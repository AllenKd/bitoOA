package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

const (
	configPath = "config"
)

type Config struct {
	Env     string
	Service struct {
		Port int
	}
	Log struct {
		Level string
	}
}

var (
	once     sync.Once
	instance *Config
)

func New() *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		instance = &Config{}
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("fail to load config file: ", err.Error())
		}
		if err := viper.Unmarshal(&instance); err != nil {
			log.Fatal(err)
		}

		log.Println("config initialized")
	})
	return instance
}
