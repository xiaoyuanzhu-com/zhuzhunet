package diagnose

import (
	"testing"

	"github.com/xiaoyuanzhu-com/zhuzhunet/cloud"
	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models/enum"
	"github.com/xiaoyuanzhu-com/zhuzhunet/utils"
)

func TestDiagnoseDNS(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	c := cloud.NewCloud(cfg.CloudURL)
	dnsList, err := c.GetDNSList()
	if err != nil {
		t.Fatal(err)
	}
	dnsList.List = append([]*models.DNS{
		{
			Address: "https://1.1.1.1/dns-query",
			Type:    enum.DNSTypeDoH,
		},
		{
			Address: "172.16.1.3",
			Type:    enum.DNSTypeUDP,
		},
		{
			Address: "202.106.46.151",
			Type:    enum.DNSTypeUDP,
		},
		{
			Address: "202.106.195.68",
			Type:    enum.DNSTypeUDP,
		},
	}, dnsList.List...)
	websiteList, err := c.GetWebsiteList()
	if err != nil {
		t.Fatal(err)
	}
	domains := []string{}
	for _, website := range websiteList.List {
		domains = append(domains, utils.GetDomain(website.Address))
	}
	for _, dns := range dnsList.List {
		diagnose := NewDiagnose(c)
		report := diagnose.queryDNS(dns, domains)
		if report.Status == enum.TestStatusFailed {
			t.Logf("dns: %s, error: %s", dns.Address, report.Message)
			continue
		}
		t.Logf("dns: %s, time: %d ms", dns.Address, report.End.Sub(report.Start).Milliseconds())
		utils.PrintJSON(report)
		break
	}
}
