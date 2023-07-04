package config

import (
	"fmt"

	"github.com/spf13/viper"
)


func GetJSONConfig() {
	viper.AddConfigPath("./")
    viper.SetConfigName("dev.config") // Register config file name (no extension)
    viper.SetConfigType("json")   // Look for specific type
    viper.ReadInConfig()

    port := viper.Get("server.port")

    fmt.Println(port)
}