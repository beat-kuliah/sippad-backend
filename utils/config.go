package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBdriver       string `mapstructure:"DB_Driver"`
	DB_source      string `mapstructure:"DB_Source"`
	DB_source_live string `mapstructure:"DB_Source_LIVE"`
}

func LoadConfig(path string) (config *Config, error error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
