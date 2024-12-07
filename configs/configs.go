package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/xiaoyuanzhu-com/zhuzhunet/logs"
	"go.uber.org/zap"
)

type Configs struct {
	LogLevel string `json:"log_level"`
	CloudURL string `json:"cloud_url"`
}

var configDir string

func isInTest() bool {
	return strings.HasSuffix(os.Args[0], ".test") || strings.Contains(os.Args[0], "debug_bin")
}

func init() {
	configDir = os.Getenv("CONFIG_DIR")
	if len(configDir) == 0 {
		if isInTest() {
			_, filename, _, _ := runtime.Caller(0)
			configDir = filepath.Dir(filename)
		} else {
			configDir = "./configs"
		}
	}
}

func Load() (*Configs, error) {
	var configs Configs
	if err := readConfig(filepath.Join(configDir, "configs.json"), &configs); err != nil {
		return nil, err
	}
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
