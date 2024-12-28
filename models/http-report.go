package models

import (
	"time"

	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
)

type HTTPReport struct {
	URL             string          `json:"url"`
	IP              string          `json:"ip"`
	Start           time.Time       `json:"start"`
	End             time.Time       `json:"end"`
	Status          enum.TestStatus `json:"status"`
	Message         string          `json:"message"`
	DNSStart        time.Time       `json:"dns_start"`
	DNSEnd          time.Time       `json:"dns_end"`
	ConnectionStart time.Time       `json:"connection_start"`
	ConnectionEnd   time.Time       `json:"connection_end"`
	TLSStart        time.Time       `json:"tls_start"`
	TLSEnd          time.Time       `json:"tls_end"`
}
