package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	logConfig        LogConfig
	httpServerConfig HTTPServerConfig
	csvPath          string
}

func (config Config) GetLogConfig() LogConfig {
	return config.logConfig
}

func (config Config) GetHTTPServerConfig() HTTPServerConfig {
	return config.httpServerConfig
}

func (config Config) GetCSVPath() string {
	return config.csvPath
}

func NewConfig(configFile string) Config {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	return Config{
		logConfig:        newLogConfig(),
		httpServerConfig: newHTTPServerConfig(),
	}
}
