package server

import (
	"encoding/json"
	"net/http"
	"net/url"

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

func (c *Cloud) GetBrands() ([]*models.Brand, error) {
	var brands []*models.Brand
	if err := c.getJSON("api/brands", &brands); err != nil {
		return nil, err
	}
	return brands, nil
}
