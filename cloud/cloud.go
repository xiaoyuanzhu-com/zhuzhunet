package cloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/xiaoyuanzhu-com/zhuzhunet/models"
)

type Cloud struct {
	baseURL string
}

func NewCloud(baseURL string) *Cloud {
	return &Cloud{baseURL: baseURL}
}

func (c *Cloud) getJSON(path string, dest interface{}) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	url := u.String()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dest)
}

func (c *Cloud) GetManifest() (*models.Manifest, error) {
	var manifest models.Manifest
	if err := c.getJSON("api/manifest", &manifest); err != nil {
		return nil, err
	}
	return &manifest, nil
}

func (c *Cloud) GetBrandList() (*models.BrandList, error) {
	var brandList models.BrandList
	if err := c.getJSON("/api/brands", &brandList); err != nil {
		return nil, err
	}
	return &brandList, nil
}

func (c *Cloud) GetDNSList() (*models.DNSList, error) {
	var dnsList models.DNSList
	if err := c.getJSON("/api/dns", &dnsList); err != nil {
		return nil, err
	}
	return &dnsList, nil
}

func (c *Cloud) GetWebsiteList() (*models.WebsiteList, error) {
	var websiteList models.WebsiteList
	if err := c.getJSON("/api/websites", &websiteList); err != nil {
		return nil, err
	}
	return &websiteList, nil
}

func (c *Cloud) GetIPInfo(ips []string) ([]*models.IPInfo, error) {
	var ipInfo []*models.IPInfo
	if err := c.getJSON(fmt.Sprintf("/api/ip/%s", strings.Join(ips, ",")), &ipInfo); err != nil {
		return nil, err
	}
	return ipInfo, nil
}
