package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
)

type Application struct {
	DBModel *sql.DB
	L       hclog.Logger
}

type config struct {
	DB struct {
		DSN string
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("No config file found")
		} else {
			panic(err.Error())
		}
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return C
}
