package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server Server `yaml:"server"`
		DB     DB     `yaml:"db"`
	}

	Server struct {
		Port int `yaml:"port"`
	}

	DB struct {
		Type            string `yaml:"type"`
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		User            string `yaml:"user"`
		Password        string `yaml:"password"`
		DBName          string `yaml:"dbname"`
		SSLMode         string `yaml:"sslmode"`
		Timezone        string `yaml:"timezone"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifetime int    `yaml:"connMaxLifetime"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		ConnMaxIdleTime int    `yaml:"connMaxIdleTime"`
	}
)

func Get() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return config, fmt.Errorf("config not found")
		} else {
			return config, fmt.Errorf("failed to read config")
		}
	}

	if err := viper.Unmarshal(config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config")
	}

	return config, nil
}
