package server

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
	"github.com/xiaoyuanzhu-com/zhuzhunet/utils"
	"go.uber.org/zap"
)

func (s *Server) GetAPIInput(c *gin.Context) (*models.APIInput, error) {
	input := &models.APIInput{
		ClientIP: c.ClientIP(),
		Path:     c.Request.URL.Path,
		Method:   c.Request.Method,
		Headers:  c.Request.Header,
		Query:    c.Request.URL.Query(),
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	input.Body = string(body)
	return input, nil
}

func (s *Server) ServeAsync(c *gin.Context, f func(record *AsyncAPIRecord)) {
	async, _ := c.GetQuery("async")
	input, err := s.GetAPIInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := utils.Hash(input)
	s.asyncMapLock.RLock()
	record, ok := s.asyncMap[hash]
	s.asyncMapLock.RUnlock()
	if !ok {
		logs.Info("async record not found, create new record", zap.String("hash", hash))
		record = &AsyncAPIRecord{
			Input: input,
			Output: &models.APIOutput{
				Data: nil,
				Meta: &models.APIOutputMeta{
					Status:  enum.APIStatusInProgress,
					Message: "",
				},
			},
		}
		s.asyncMapLock.Lock()
		s.asyncMap[hash] = record
		s.asyncMapLock.Unlock()
		if async == "true" {
			go f(record)
		} else {
			f(record)
		}
	}
	c.JSON(http.StatusOK, record.Output)
	if record.Output.Meta.Status == enum.APIStatusSuccess || record.Output.Meta.Status == enum.APIStatusError {
		s.asyncMapLock.Lock()
		delete(s.asyncMap, hash)
		s.asyncMapLock.Unlock()
	}
}
