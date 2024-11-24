package base

import (
	"context"

	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
)

type ServerContext struct {
	Context context.Context
	Configs *configs.Configs
}
