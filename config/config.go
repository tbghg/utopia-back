package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var V *viper.Viper

func init() {
	V = viper.New()
	V.SetConfigName("config")
	V.AddConfigPath("config")
	err := V.ReadInConfig()
	if err != nil {
		fmt.Printf("read config failed: %v\n", err)
	}
}
