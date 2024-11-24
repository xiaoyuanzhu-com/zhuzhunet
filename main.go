package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/server"
	"go.uber.org/zap"
)

func waitSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}

func main() {
	var srv *server.Server
	configs.OnLoadOrChange(func(cfg *configs.Configs) {
		if srv != nil {
			if err := srv.Stop(); err != nil {
				logs.Error("stop server error", zap.Error(err))
				return
			}
			srv = nil
		}
		srv = server.NewServer(cfg)
		if err := srv.Start(); err != nil {
			logs.Error("start server error", zap.Error(err))
			return
		}
	})
	waitSignal()
	if srv != nil {
		if err := srv.Stop(); err != nil {
			logs.Error("stop server error", zap.Error(err))
		}
	}
}
