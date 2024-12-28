package server

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
)

type PingQuery struct {
	Address string `json:"address" form:"address" query:"address"`
	Count   int    `json:"count" form:"count" query:"count"`
}

func (s *Server) DiagnosePingHandler(c *gin.Context) {
	s.ServeAsync(c, func(record *AsyncAPIRecord) {
		query := PingQuery{}
		if err := c.ShouldBindQuery(&query); err != nil {
			record.Output.Meta.Status = enum.APIStatusError
			record.Output.Meta.Message = err.Error()
			return
		}
		var wg sync.WaitGroup
		wg.Add(1)
		s.diagnose.Ping(query.Address, query.Count, func(report *models.PingReport) {
			record.Output.Data = report
		}, func(report *models.PingReport) {
			record.Output.Data = report
			record.Output.Meta.Status = enum.APIStatusSuccess
			record.Output.Meta.Message = "ping success"
			wg.Done()
		})
		wg.Wait()
	})
}
