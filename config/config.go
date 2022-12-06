package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	viperconfig := viper.New()
	viperconfig.AddConfigPath(".")
	viperconfig.SetConfigName("config")
	viperconfig.SetConfigType("yaml")
	if err := viperconfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found", err)
		} else {
			fmt.Println("read config err", err)
		}
	}

	LoadServer(viperconfig)
	LoadData(viperconfig)

}

func LoadServer(viperconfig *viper.Viper) {
	AppMode = viperconfig.GetString("server.AppMode")
	HttpPort = viperconfig.GetString("server.HttpPort")
}

func LoadData(viperconfig *viper.Viper) {
	Db = viperconfig.GetString("database.Db")
	DbHost = viperconfig.GetString("database.DbHost")
	DbPort = viperconfig.GetString("database.DbPort")
	DbUser = viperconfig.GetString("database.DbUser")
	DbPassWord = viperconfig.GetString("database.DbPassWord")
	DbName = viperconfig.GetString("database.DbName")
}
