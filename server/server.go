package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/xiaoyuanzhu-com/zhuzhunet/cloud"
	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/diagnose"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"go.uber.org/zap"
)

type AsyncAPIRecord struct {
	Input  *models.APIInput
	Output *models.APIOutput
}

type Server struct {
	configs    *configs.Configs
	httpServer *http.Server
	cloud      *cloud.Cloud
	diagnose   *diagnose.Diagnose

	asyncMap     map[string]*AsyncAPIRecord
	asyncMapLock sync.RWMutex
}

func NewServer(configs *configs.Configs) *Server {
	return &Server{
		configs: configs,
	}
}

func (s *Server) Start() error {
	logs.Info("start server", zap.Any("configs", s.configs))
	s.cloud = cloud.NewCloud(s.configs.CloudURL)
	s.diagnose = diagnose.NewDiagnose(s.cloud)
	s.asyncMap = make(map[string]*AsyncAPIRecord)
	if err := s.startAPI(); err != nil {
		return err
	}
	return nil
}

func (s *Server) startAPI() error {
	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(ginzap.Ginzap(logs.GetLogger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logs.GetLogger(), true))
	s.RegisterRoutes(r)
	r.NoRoute(static.Serve("/", static.LocalFile("./ui/dist", false)))
	s.httpServer = &http.Server{
		Addr:    ":27831",
		Handler: r,
	}
	var err error
	go func() {
		err = s.httpServer.ListenAndServe()
	}()
	time.Sleep(5 * time.Millisecond)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	logs.Info("stop server")
	if err := s.stopAPI(); err != nil {
		return err
	}
	return nil
}

func (s *Server) stopAPI() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
