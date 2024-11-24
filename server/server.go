package server

import (
	"github.com/xiaoyuanzhu-com/zhuzhunet/base"
	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"go.uber.org/zap"
)

type Server struct {
	ctx *base.ServerContext
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
	return nil
}

func (s *Server) Stop() error {
	logs.Info("stop server")
	return nil
}
