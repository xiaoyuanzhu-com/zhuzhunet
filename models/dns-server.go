package models

import "github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"

type DNSServer struct {
	ID        string             `json:"id" bson:"id"`
	Address   string             `json:"address" bson:"address"`
	Type      enum.DNSServerType `json:"type" bson:"type"`
	BrandName string             `json:"brand_name" bson:"brand_name"`
	Desc      string             `json:"desc" bson:"desc"`
}
