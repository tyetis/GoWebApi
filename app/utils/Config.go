package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}
	return config
}
