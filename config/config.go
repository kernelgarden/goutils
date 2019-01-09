package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Read(configPath, configName string, config interface{}) error {
	viper.SetConfigName(configName)
	setConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error in config file: %s\n", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Fatal error in config file: %s\n", err)
	}

	return nil
}

func setConfigPath(path string) {
	viper.AddConfigPath(path)
}