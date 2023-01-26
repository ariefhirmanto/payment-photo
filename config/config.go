package config

import (
	"log"

	_viper "github.com/spf13/viper"
)

type MainConfig struct {
	Database DatabaseConfig
	Midtrans PaymentConfig
	Server   ServerConfig
	Product  ProductConfig
}

type DatabaseConfig struct {
	Host   string
	Port   string
	DBName string
	DBUser string
	DBPass string
}

type PaymentConfig struct {
	ClientKey string
	ServerKey string
	APIEnv    string
}

type ServerConfig struct {
	Port string
}

type ProductConfig struct {
	LocalEvent  bool
	AdminSwitch bool
}

func LoadConfig() (config MainConfig) {
	_viper.AddConfigPath("/app/config")
	_viper.AddConfigPath("./config")
	_viper.SetConfigType("yaml")
	_viper.SetConfigName("config.local") // read .yaml config
	err := _viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	_viper.UnmarshalKey("Database", &config.Database)
	_viper.UnmarshalKey("Midtrans", &config.Midtrans)
	_viper.UnmarshalKey("Server", &config.Server)
	_viper.UnmarshalKey("Product", &config.Product)
	return
}
