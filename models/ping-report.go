package models

import (
	"time"

	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
)

type PingReport struct {
	Address               string          `json:"address"`
	IP                    string          `json:"ip"`
	Start                 time.Time       `json:"start"`
	End                   time.Time       `json:"end"`
	Status                enum.TestStatus `json:"status"`
	Message               string          `json:"message"`
	PacketsSent           int             `json:"packets_sent"`
	PacketsRecv           int             `json:"packets_recv"`
	PacketsRecvDuplicates int             `json:"packets_recv_duplicates"`
	PacketLoss            float64         `json:"packet_loss"`
	RTTNs                 []int64         `json:"rtt_ns"`
	RTTTimestamps         []time.Time     `json:"rtt_timestamps"`
	MinRttNs              int64           `json:"min_rtt_ns"`
	MaxRttNs              int64           `json:"max_rtt_ns"`
	AvgRttNs              int64           `json:"avg_rtt_ns"`
	StdDevRttNs           int64           `json:"std_dev_rtt_ns"`
}
