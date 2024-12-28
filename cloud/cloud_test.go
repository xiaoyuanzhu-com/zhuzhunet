package cloud

import (
	"testing"

	"github.com/xiaoyuanzhu-com/zhuzhunet/configs"
	"github.com/xiaoyuanzhu-com/zhuzhunet/utils"
)

func TestGetManifest(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	cloud := NewCloud(cfg.CloudURL)
	manifest, err := cloud.GetManifest()
	if err != nil {
		t.Fatal(err)
	}
	utils.PrintJSON(manifest)
}

func TestGetDNSList(t *testing.T) {
	cfg, err := configs.Load()
	if err != nil {
		t.Fatal(err)
	}
	cloud := NewCloud(cfg.CloudURL)
	dnsList, err := cloud.GetDNSList()
	if err != nil {
		t.Fatal(err)
	}
	utils.PrintJSON(dnsList)
}
