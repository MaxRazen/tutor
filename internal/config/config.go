package config

import "strings"

type RuntimeConfig struct {
	Host    string
	Port    string
	URL     string
	DevMode bool
}

func (rc *RuntimeConfig) GetServerHost() string {
	return rc.Host + ":" + rc.Port
}

func NewConfig(mode string, args []string) RuntimeConfig {
	return RuntimeConfig{
		Host:    getArg("--host", args, ""),
		Port:    getArg("--port", args, "3000"),
		URL:     getArg("--url", args, "http://localhost:3000"),
		DevMode: mode == "devonly",
	}
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
	}

	return d
}
