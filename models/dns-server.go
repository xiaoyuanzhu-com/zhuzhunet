package models

import "github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"

type DNSServer struct {
	ID      string             `json:"id" bson:"id"`
	Type    enum.DNSServerType `json:"type" bson:"type"`
	Address string             `json:"address" bson:"address"`
	Name    string             `json:"name" bson:"name"`
	Logo    string             `json:"logo" bson:"logo"`
}
