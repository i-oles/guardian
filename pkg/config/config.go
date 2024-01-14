package config

import (
	"flag"
	"github.com/tkanos/gonfig"
	"log/slog"
	"path/filepath"
)

const (
	defaultConfigName = "dev"
	ext               = ".json"
)

func GetConfig(cfgBasePath string, cfgAddr any) error {
	profile := flag.String("profile", "", "config name to load from")
	flag.Parse()

	configName := defaultConfigName
	if *profile != "" {
		configName = *profile
	}

	cfgPath := filepath.Join(cfgBasePath, configName+ext)

	err := gonfig.GetConf(cfgPath, cfgAddr)
	if err != nil {
		slog.Error("could not get config file %v", err)
	}

	return nil
}
