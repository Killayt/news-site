package config

import (
	"log"

	"github.com/spf13/viper"
)

type Conf struct {
	Port     string
	Database Database
}

type Database struct {
	DBName     string
	DBLogin    string
	DBPassword string
}

type ConfInterface interface {
	LoadConf()
}

func LoadConf(pathToConf string) (*Conf, error) {

	v := viper.New()

	v.SetConfigFile(pathToConf)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Config file is empty")
		return nil, err
	}

	return &Conf{
		Port: v.GetString("Port"),
		Database: Database{
			DBName:     v.GetString("DBname"),
			DBLogin:    v.GetString("DBLogin"),
			DBPassword: v.GetString("DBPassword"),
		},
	}, nil
}
