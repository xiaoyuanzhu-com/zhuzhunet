package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/xiaoyuanzhu-com/zhuzhunet/base"
	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"go.uber.org/zap"
)

type Server struct {
	ctx *base.ServerContext

	httpServer *http.Server
}

func NewServer(configs *configs.Configs) *Server {
	return &Server{
		ctx: &base.ServerContext{
			Configs: configs,
		},
	}
}

func (s *Server) Start() error {
	logs.Info("start server", zap.Any("configs", s.ctx.Configs))
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
