package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Mongo *MongoConfiguration `json:"mongo"`
	Flags *Flags
}

type MongoConfiguration struct {
	URI          string `json:"uri"`
	DatabaseName string `json:"database_name"`
}

func ReadConfiguration(f *Flags) (*Configuration, error) {
	var config = &Configuration{}
	viper.SetConfigName(f.Environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.Flags = f

	return config, nil
}

func SplitConfigs(c *Configuration) *MongoConfiguration {
	return c.Mongo
}
