package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN     string `mapstructure:"DB_URL"`
	PORT    string `mapstructure:"PORT"`
	UI_PATH string `mapstructure:"UI_PATH"`
}

func LoadConfig(path string, name string, ext string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(ext)

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("viper read config failed: ", err)
		return
	}

	err = viper.Unmarshal(&config)

	return

}
