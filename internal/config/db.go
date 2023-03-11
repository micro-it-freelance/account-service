package config

import "github.com/spf13/viper"

var DB struct {
	Name     string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
}

func init() {
	DB.Name = viper.GetString("DB_NAME")
	DB.User = viper.GetString("DB_USER")
	DB.Password = viper.GetString("DB_PASSWORD")
	DB.Host = viper.GetString("DB_HOST")
	DB.Port = viper.GetInt("DB_PORT")
}
