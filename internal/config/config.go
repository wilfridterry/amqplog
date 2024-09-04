package config

import (

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Enviroment string

	DB Mongo

	Rabbitmq Rabbitmq
}

type Mongo struct {
	URI      string
	Database string
	Username string
	Password string
}

type Rabbitmq struct {
	Host     string
	Port     uint16
	Queue    string
	Username string
	Password string
}

func NewConfig() (*Config, error) {
	cf := new(Config)

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := envconfig.Process("db", &cf.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("rabbitmq", &cf.Rabbitmq); err != nil {
		return nil, err
	}

	return cf, nil
}
