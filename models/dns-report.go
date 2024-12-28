package models

import (
	"time"

	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
)

type DNSReport struct {
	DNS     *DNS               `json:"dns"`
	Start   time.Time          `json:"start"`
	End     time.Time          `json:"end"`
	Status  enum.TestStatus    `json:"status"`
	Message string             `json:"message"`
	Records []*DNSReportRecord `json:"records"`
}

type DNSReportRecord struct {
	Domain   string    `json:"domain"`
	IPs      []*IPInfo `json:"ips"`
	Duration int64     `json:"duration"`
}
