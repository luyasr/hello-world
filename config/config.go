package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Port int `json:"port"`
}

var Conf = new(Config)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("fatal error config file: %s", err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		log.Panicf("unmarshal conf failed, err: %s", err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			log.Panicf("unmarshal conf failed, err: %s", err)
		}
	})
	viper.WatchConfig()
}
