package diagnose

import (
	"crypto/tls"
	"net/http/httptrace"
	"net/textproto"

	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/utils"
)

func (d *Diagnose) DiagnoseHTTP(u string, ip string) *models.HTTPReport {
	// report := &models.HTTPReport{
	// 	URL:    u,
	// 	IP:     ip,
	// 	Start:  time.Now(),
	// 	Status: enum.TestStatusRunning,
	// }
	trace := &httptrace.ClientTrace{
		GetConn:              func(hostPort string) {},
		GotConn:              func(info httptrace.GotConnInfo) {},
		PutIdleConn:          func(err error) {},
		GotFirstResponseByte: func() {},
		Got100Continue:       func() {},
		Got1xxResponse:       func(code int, header textproto.MIMEHeader) error { return nil },
		DNSStart:             func(info httptrace.DNSStartInfo) {},
		DNSDone:              func(info httptrace.DNSDoneInfo) {},
		ConnectStart:         func(network, addr string) {},
		ConnectDone:          func(network, addr string, err error) {},
		TLSHandshakeStart:    func() {},
		TLSHandshakeDone:     func(state tls.ConnectionState, err error) {},
		WroteHeaderField:     func(key string, values []string) {},
		WroteHeaders:         func() {},
		Wait100Continue:      func() {},
		WroteRequest:         func(info httptrace.WroteRequestInfo) {},
	}

	utils.PrintJSON(trace)
	return nil
}
