package diagnose

import (
	"errors"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
)

func (d *Diagnose) Ping(address string, count int, onUpdate func(report *models.PingReport), onFinish func(report *models.PingReport)) {
	report := &models.PingReport{
		Address:       address,
		Start:         time.Now(),
		Status:        enum.TestStatusRunning,
		RTTTimestamps: []time.Time{},
	}
	pinger, err := probing.NewPinger(address)
	if err != nil {
		report.Status = enum.TestStatusFailed
		report.Message = err.Error()
		report.End = time.Now()
		onFinish(report)
		return
	}
	pinger.Count = count
	updateReport := func() {
		stats := pinger.Statistics()
		report.IP = stats.IPAddr.IP.String()
		report.PacketsSent = stats.PacketsSent
		report.PacketsRecv = stats.PacketsRecv
		report.PacketsRecvDuplicates = stats.PacketsRecvDuplicates
		report.PacketLoss = stats.PacketLoss
		report.RTTNs = []int64{}
		for _, rtt := range stats.Rtts {
			report.RTTNs = append(report.RTTNs, rtt.Nanoseconds())
		}
		for {
			if len(report.RTTTimestamps) >= len(report.RTTNs) {
				break
			}
			report.RTTTimestamps = append(report.RTTTimestamps, time.Now())
		}
		report.MinRttNs = stats.MinRtt.Nanoseconds()
		report.MaxRttNs = stats.MaxRtt.Nanoseconds()
		report.AvgRttNs = stats.AvgRtt.Nanoseconds()
		report.StdDevRttNs = stats.StdDevRtt.Nanoseconds()
	}
	var errs error
	timeout := time.Second * 5
	// don't set pinger.Timeout, it's the total timeout for any count, e.g. we need to ping infinite times
	// pinger.Timeout = timeout
	ticker := time.NewTicker(timeout)
	go func() {
		<-ticker.C
		errs = errors.Join(errs, errors.New("ping timeout"))
		pinger.Stop()
	}()
	pinger.OnRecv = func(p *probing.Packet) {
		updateReport()
		onUpdate(report)
		ticker.Reset(timeout)
	}
	pinger.OnSendError = func(p *probing.Packet, e error) {
		// logs.Warn("ping OnSendError", zap.Error(e))
		// report.Status = enum.TestStatusFailed
		// report.Message = e.Error()
		// report.End = time.Now()
		// errs = errors.Join(errs, e)
		// pinger.Stop()
	}
	// https://github.com/prometheus-community/pro-bing/issues/105
	pinger.OnRecvError = func(e error) {
		// logs.Warn("ping OnRecvError", zap.Error(e))
		// report.Status = enum.TestStatusFailed
		// report.Message = e.Error()
		// report.End = time.Now()
		// errs = errors.Join(errs, e)
		// pinger.Stop()
	}
	err = pinger.Run()
	errs = errors.Join(errs, err)
	if errs != nil {
		report.Status = enum.TestStatusFailed
		report.Message = errs.Error()
		report.End = time.Now()
		onFinish(report)
		return
	}
	updateReport()
	report.Status = enum.TestStatusCompleted
	report.End = time.Now()
	onFinish(report)
}
