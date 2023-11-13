package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var C = new(Config)

type Config struct {
	HttpServer HttpServer `json:"http_server"`
	GrpcServer GrpcServer `json:"grpc_server"`
}

type HttpServer struct {
	Port int `json:"port"`
}

type GrpcServer struct {
	Port int `json:"port"`
}

func (h *HttpServer) Addr() string {
	return fmt.Sprintf(":%d", h.Port)
}

func (g *GrpcServer) Addr() string {
	return fmt.Sprintf(":%d", g.Port)
}

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
	if err := viper.Unmarshal(C); err != nil {
		log.Panicf("unmarshal conf failed, err: %s", err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(C); err != nil {
			log.Panicf("unmarshal conf failed, err: %s", err)
		}
	})
	viper.WatchConfig()
}
