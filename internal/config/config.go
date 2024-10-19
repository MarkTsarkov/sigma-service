package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	Port int    `envconfig:"PORT"`
	DB   string `envconfig:"DB"`
}

type Environment interface {
	GetPort() int
	GetDB() string
}

func (e *environment) GetPort() int {
	return e.Port
}

func (e *environment) GetDB() string {
	return e.DB
}

func NewEnv() (Environment, error) {
	var e environment

	err := envconfig.Process("myapp", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &e, nil
}
