package config

import "sync"

type Config struct {
	AppName string
	Version string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			AppName: "NotificationService",
			Version: "1.0.0",
		}
	})
	return instance
}
