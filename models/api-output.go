package models

import "github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"

type APIOutput struct {
	Data interface{}    `json:"data"`
	Meta *APIOutputMeta `json:"meta"`
}

type APIOutputMeta struct {
	Status  enum.APIStatus `json:"status"`
	Message string         `json:"message"`
}
