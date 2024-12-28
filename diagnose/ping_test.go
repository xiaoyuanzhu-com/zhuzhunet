package diagnose

import (
	"testing"

	"github.com/xiaoyuanzhu-com/zhuzhunet/cloud"
	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
	"github.com/xiaoyuanzhu-com/zhuzhunet/utils"
)

func TestDiagnosePing(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	c := cloud.NewCloud(cfg.CloudURL)
	d := NewDiagnose(c)
	d.Ping("www.baidu.com", 4, func(report *models.PingReport) {
	}, func(report *models.PingReport) {
		utils.PrintJSON(report)
	})
}

func TestDiagnosePingInfinite(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	c := cloud.NewCloud(cfg.CloudURL)
	d := NewDiagnose(c)
	d.Ping("www.baidu.com", 0, func(report *models.PingReport) {
		utils.PrintJSON(report)
	}, func(report *models.PingReport) {
		utils.PrintJSON(report)
	})
}

func TestDiagnosePingError(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	c := cloud.NewCloud(cfg.CloudURL)
	d := NewDiagnose(c)
	d.Ping("www.some-domain-that-does-not-exist.com", 4, func(report *models.PingReport) {
	}, func(report *models.PingReport) {
		utils.PrintJSON(report)
	})
}

func TestDiagnosePingTimeout(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	c := cloud.NewCloud(cfg.CloudURL)
	d := NewDiagnose(c)
	d.Ping("www.google.com", 4, func(report *models.PingReport) {
		utils.PrintJSON(report)
	}, func(report *models.PingReport) {
		utils.PrintJSON(report)
	})
}
