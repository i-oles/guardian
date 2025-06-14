package config

import "encoding/json"

type Configuration struct {
	DBFileName   string
	BulbCollName string
	MockBulbs    bool
	Logging      bool
}

func (c *Configuration) Pretty() string {
	cfgPretty, _ := json.MarshalIndent(c, "", "  ")

	return string(cfgPretty)
}
