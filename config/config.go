package config

import (
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Providers is a Wire provider set that provides a new Config.
var Providers = wire.NewSet(NewConfig)

// Config holds all configuration for the application.
type Config struct {
	App       AppConfig      `mapstructure:"app"`
	Database  DatabaseConfig `mapstructure:"database"`
	Redis     RedisConfig    `mapstructure:"redis"`
	JWT       JWTConfig      `mapstructure:"jwt"`
	Signature Signature      `mapstructure:"signature"`
}

type Signature struct {
	Secret string `mapstructure:"signature.secret"`
}

// AppConfig holds application-specific configuration.
type AppConfig struct {
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

// DatabaseConfig holds database configuration.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// DSN returns the Data Source Name for connecting to the database.
func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.Name)
}

// RedisConfig holds redis configuration.
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// JWTConfig holds JWT configuration.
type JWTConfig struct {
	Secret string `mapstructure:"jwt.secret"`
}

// NewConfig creates a new Config instance.
func NewConfig() (*Config, error) {
	viper.SetConfigFile("./config/config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
