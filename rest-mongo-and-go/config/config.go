package config

import (
	viper "github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

func GetConfig() Configuration {
	conf := Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("yml")
	err := viper.ReadConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}
