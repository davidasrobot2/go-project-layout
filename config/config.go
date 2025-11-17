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
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type Signature struct {
	Secret string `mapstructure:"secret"`
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
	Password string `mapstructure:"passowrd"` // Note: "passowrd" matches the typo in your config.yaml
	Name     string `mapstructure:"name"`
}

// DSN returns the Data Source Name for connecting to the database.
func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		d.Host, d.User, d.Password, d.Name, d.Port)
}

// RedisConfig holds redis configuration.
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// JWTConfig holds JWT configuration.
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

// NewConfig creates a new Config instance.
func NewConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
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
