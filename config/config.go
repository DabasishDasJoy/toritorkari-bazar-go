package config

import (
	"log"

	"github.com/spf13/viper"
)

// mapping env variables
type Config struct {
	DBUser string `mapstructure:"DBUSER"`
	DBPass string `mapstructure:"DBPASS"`
	DBIp   string `mapstructure:"DBIP"`
	DBName string `mapstructure:"DBNAME"`
	Port   string `mapstructure:"PORT"`
	JWTKEY string `mapstructure:"JWTKEY"`
	CERT   string `mapstructure:"CERT"`
}

func InitConfig() *Config {
	// viper.AddConfigPath(".")
	// viper.SetConfigName("app")
	// viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return config
}

var LocalConfig *Config

func SetConfig() {
	LocalConfig = InitConfig()
}
