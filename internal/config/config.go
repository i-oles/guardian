package config

import "encoding/json"

type Configuration struct {
	DBFileName string
	MockBulbs  bool
	IsDebugOn  bool
}

func (c *Configuration) Pretty() string {
	cfgPretty, _ := json.MarshalIndent(c, "", "  ")

	return string(cfgPretty)
}
