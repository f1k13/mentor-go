package config

import (
	"server/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BingIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

var instance *Config
var once *sync.Once

func GetConfig() *Config {
	once.Do(func() {
		Logger := logging.GetLogger()
		Logger.Info("Loading config file")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			Logger.Info(help)
			Logger.Fatal(err)
		}

	})
	return instance
}
