package config

import "strings"

const (
	FlagPresent = "_flagPresent_"
)

type RuntimeConfig struct {
	Host        string
	Port        string
	BaseUrl     string
	DevMode     bool
	AutoMigrate bool
}

var runtimeConfig *RuntimeConfig

func (rc *RuntimeConfig) GetServerHost() string {
	return rc.Host + ":" + rc.Port
}

func NewConfig(mode string, args []string) *RuntimeConfig {
	runtimeConfig = &RuntimeConfig{
		Host:        getArg("--host", args, ""),
		Port:        getArg("--port", args, "3000"),
		BaseUrl:     getArg("--base-url", args, "http://localhost:3000"),
		DevMode:     mode == "devonly",
		AutoMigrate: getArg("--auto-migrate", args, "") == FlagPresent,
	}
	return runtimeConfig
}

func GetRuntimeConfig() *RuntimeConfig {
	return runtimeConfig
}

func getArg(key string, args []string, d string) string {
	for _, v := range args {
		if !strings.HasPrefix(v, key) {
			continue
		}

		t := strings.Split(v, "=")

		if len(t) > 1 {
			return t[1]
		}
		return FlagPresent
	}

	return d
}
