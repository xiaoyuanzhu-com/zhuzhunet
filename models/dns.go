package models

import "github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"

type DNS struct {
	ID      int          `json:"id" noco:"Id"`
	Address string       `json:"address" noco:"address"`
	Type    enum.DNSType `json:"type" noco:"type"`
	Desc    string       `json:"desc" noco:"desc"`
	Brands  struct {
		ID   int    `json:"id" noco:"Id"`
		Name string `json:"name" noco:"name"`
	} `json:"brands" noco:"brands"`
}
