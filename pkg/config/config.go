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
	cfgFlag := flag.String("cfg", "", "config name to load from")
	flag.Parse()

	configName := defaultConfigName
	if *cfgFlag != "" {
		configName = *cfgFlag
	}

	cfgPath := filepath.Join(cfgBasePath, configName+ext)

	err := gonfig.GetConf(cfgPath, cfgAddr)
	if err != nil {
		slog.Error("could not get config file %v", err)
	}

	return nil
}
