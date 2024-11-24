package configs

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"go.uber.org/zap"
)

type Configs struct {
	LogLevel string     `json:"log_level"`
	DNS      DNSConfigs `json:"dns"`
}

type DNSConfigs struct {
	Servers []DNSServer `json:"servers"`
}

type DNSServerType string

const (
	DNSServerTypeDefault  DNSServerType = "default"
	DNSServerTypeUDP      DNSServerType = "udp"
	DNSServerTypeDoH      DNSServerType = "doh"
	DNSServerTypeDoT      DNSServerType = "dot"
	DNSServerTypeDoQ      DNSServerType = "doq"
	DNSServerTypeDNSCrypt DNSServerType = "dnscrypt"
)

type DNSServer struct {
	Name string        `json:"name"`
	Type DNSServerType `json:"type"`
	Addr string        `json:"addr"`
	Logo string        `json:"logo"`
}

var configDir string

func init() {
	configDir = os.Getenv("CONFIG")
	if len(configDir) == 0 {
		configDir = "/config"
	}
}

func Load() (*Configs, error) {
	var dnsConfigs DNSConfigs
	if err := readConfig(filepath.Join(configDir, "dns.json"), &dnsConfigs); err != nil {
		return nil, err
	}
	var configs Configs
	if err := readConfig(filepath.Join(configDir, "configs.json"), &configs); err != nil {
		return nil, err
	}
	configs.DNS = dnsConfigs
	return &configs, nil
}

func OnLoadOrChange(callback func(configs *Configs)) {
	cfg, err := Load()
	if err == nil {
		callback(cfg)
	} else {
		logs.Error("load configs error", zap.Error(err))
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logs.Error("create watcher error", zap.Error(err))
		return
	}

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					cfg, err := Load()
					if err == nil {
						callback(cfg)
					} else {
						logs.Error("load configs error", zap.Error(err))
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logs.Error("watcher error", zap.Error(err))
			}
		}
	}()

	err = watcher.Add(configDir)
	if err != nil {
		logs.Error("add watcher error", zap.Error(err))
		return
	}
}

func readConfig(path string, v interface{}) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
