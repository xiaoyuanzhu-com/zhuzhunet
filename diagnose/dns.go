package diagnose

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/phuslu/fastdns"
	"github.com/xiaoyuanzhu-com/zhuzhunet/cloud"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
)

type Diagnose struct {
	cloud *cloud.Cloud
}

func NewDiagnose(cloud *cloud.Cloud) *Diagnose {
	return &Diagnose{
		cloud: cloud,
	}
}

func (d *Diagnose) callWithRetry(f func() error, retry int) error {
	var errs error
	for i := 0; i < retry; i++ {
		err := f()
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			return nil
		}
	}
	return errs
}

func (d *Diagnose) queryUDPSingle(dns_ *models.DNS, network string, domain string) ([]net.IP, error) {
	msg := new(dns.Msg)
	if network == "ip4" {
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	} else if network == "ip6" {
		msg.SetQuestion(dns.Fqdn(domain), dns.TypeAAAA)
	} else {
		return nil, errors.New("invalid network")
	}
	if !strings.Contains(dns_.Address, ":") {
		dns_.Address = net.JoinHostPort(dns_.Address, "53")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	response, err := dns.ExchangeContext(ctx, msg, dns_.Address)
	if err != nil {
		return nil, err
	}
	if response.Rcode != dns.RcodeSuccess {
		return nil, errors.New(dns.RcodeToString[response.Rcode])
	}
	ips := []net.IP{}
	for _, r := range response.Answer {
		if r.Header().Rrtype == dns.TypeA {
			ips = append(ips, r.(*dns.A).A)
		} else if r.Header().Rrtype == dns.TypeAAAA {
			ips = append(ips, r.(*dns.AAAA).AAAA)
		}
	}
	return ips, nil
}

func (d *Diagnose) queryUDP(dns *models.DNS, domains []string) *models.DNSReport {
	report := &models.DNSReport{
		DNS:     dns,
		Start:   time.Now(),
		End:     time.Now(),
		Status:  enum.TestStatusRunning,
		Message: "",
		Records: []*models.DNSReportRecord{},
	}
	for _, domain := range domains {
		for _, network := range []string{"ip4", "ip6"} {
			err := d.callWithRetry(func() error {
				t0 := time.Now()
				ips, err := d.queryUDPSingle(dns, network, domain)
				if err != nil {
					return err
				}
				duration := time.Since(t0)
				ipInfos := []*models.IPInfo{}
				for _, ip := range ips {
					ipInfos = append(ipInfos, &models.IPInfo{
						IP: ip.String(),
					})
				}
				report.Records = append(report.Records, &models.DNSReportRecord{
					Domain:   domain,
					IPs:      ipInfos,
					Duration: duration.Milliseconds(),
				})
				return nil
			}, 1)
			if err != nil {
				report.Status = enum.TestStatusFailed
				report.Message = err.Error()
				report.End = time.Now()
				return report
			}
		}
	}
	report.End = time.Now()
	report.Status = enum.TestStatusCompleted
	return report
}

func (d *Diagnose) queryDoT(dns *models.DNS, domains []string) *models.DNSReport {
	return &models.DNSReport{
		DNS:     dns,
		Start:   time.Now(),
		End:     time.Now(),
		Status:  enum.TestStatusFailed,
		Message: "not implemented",
		Records: []*models.DNSReportRecord{},
	}
}

func (d *Diagnose) queryDoH(dns *models.DNS, domains []string) *models.DNSReport {
	report := &models.DNSReport{
		DNS:     dns,
		Start:   time.Now(),
		End:     time.Now(),
		Status:  enum.TestStatusRunning,
		Message: "",
		Records: []*models.DNSReportRecord{},
	}
	endpoint, err := url.Parse(dns.Address)
	if err != nil {
		report.Status = enum.TestStatusFailed
		report.Message = err.Error()
		report.End = time.Now()
		return report
	}
	client := &fastdns.Client{
		Addr: endpoint.String(),
		Dialer: &fastdns.HTTPDialer{
			Endpoint: endpoint,
			Header: http.Header{
				"content-type": {"application/dns-message"},
			},
			Transport: &http.Transport{
				ForceAttemptHTTP2:   true,
				MaxIdleConns:        100,
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,
				TLSClientConfig: &tls.Config{
					NextProtos:         []string{"h2"},
					InsecureSkipVerify: false,
					ServerName:         endpoint.Hostname(),
					ClientSessionCache: tls.NewLRUClientSessionCache(1024),
				},
			},
		},
	}
	for _, domain := range domains {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		t0 := time.Now()
		ips, err := client.LookupNetIP(ctx, "ip", domain)
		if err != nil {
			report.Status = enum.TestStatusFailed
			report.Message = err.Error()
			report.End = time.Now()
			return report
		}
		duration := time.Since(t0)
		ipInfos := []*models.IPInfo{}
		for _, ip := range ips {
			ipInfos = append(ipInfos, &models.IPInfo{
				IP: ip.String(),
			})
		}
		report.Records = append(report.Records, &models.DNSReportRecord{
			Domain:   domain,
			IPs:      ipInfos,
			Duration: duration.Milliseconds(),
		})
	}
	report.End = time.Now()
	report.Status = enum.TestStatusCompleted
	return report
}

func (d *Diagnose) queryDNSCrypt(dns *models.DNS, domains []string) *models.DNSReport {
	return &models.DNSReport{
		DNS:     dns,
		Start:   time.Now(),
		End:     time.Now(),
		Status:  enum.TestStatusFailed,
		Message: "not implemented",
		Records: []*models.DNSReportRecord{},
	}
}

func (d *Diagnose) queryDoQ(dns *models.DNS, domains []string) *models.DNSReport {
	return &models.DNSReport{
		DNS:     dns,
		Start:   time.Now(),
		End:     time.Now(),
		Status:  enum.TestStatusFailed,
		Message: "not implemented",
		Records: []*models.DNSReportRecord{},
	}
}

func (d *Diagnose) enrichIPInfo(report *models.DNSReport) {
	ips := []string{}
	for _, record := range report.Records {
		for _, ip := range record.IPs {
			ips = append(ips, ip.IP)
		}
	}
	ipInfos, err := d.cloud.GetIPInfo(ips)
	if err != nil {
		return
	}
	ipMap := map[string]*models.IPInfo{}
	for _, ipInfo := range ipInfos {
		ipMap[ipInfo.IP] = ipInfo
	}
	for _, record := range report.Records {
		for _, ip := range record.IPs {
			if ipInfo, ok := ipMap[ip.IP]; ok {
				ip.Country = ipInfo.Country
				ip.City = ipInfo.City
				ip.Latitude = ipInfo.Latitude
				ip.Longitude = ipInfo.Longitude
				ip.ASN = ipInfo.ASN
				ip.AS = ipInfo.AS
			}
		}
	}
}

func (d *Diagnose) queryDNS(dns *models.DNS, domains []string) *models.DNSReport {
	var report *models.DNSReport
	switch dns.Type {
	case enum.DNSTypeUDP:
		report = d.queryUDP(dns, domains)
	case enum.DNSTypeDoT:
		return d.queryDoT(dns, domains)
	case enum.DNSTypeDoH:
		report = d.queryDoH(dns, domains)
	case enum.DNSTypeDNSCrypt:
		report = d.queryDNSCrypt(dns, domains)
	case enum.DNSTypeDoQ:
		return d.queryDoQ(dns, domains)
	default:
		report = &models.DNSReport{
			DNS:     dns,
			Start:   time.Now(),
			End:     time.Now(),
			Status:  enum.TestStatusFailed,
			Message: "invalid dns type",
			Records: []*models.DNSReportRecord{},
		}
	}
	if report.Status == enum.TestStatusCompleted {
		d.enrichIPInfo(report)
	}
	return report
}

func (d *Diagnose) DiagnoseDNS() {
	// d.diagnoseDNSSingle("8.8.8.8:53", "www.baidu.com")
}
