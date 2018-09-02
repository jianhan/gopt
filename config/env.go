package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("unable to load .env file", err)
		panic(err)
	}
}

type Env struct {
	Host         string `env:"HOST" envDefault:"127.0.0.1"`
	Port         int    `env:"PORT" envDefault:"8888"`
	WriteTimeout int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	ReadTimeout  int    `env:"READ_TIMEOUT" envDefault:"15"`
	IdleTimeout  int    `env:"IDLE_TIMEOUT" envDefault:"15"`
}

func (e *Env) Address() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

func EnvConfigs() (*Env, error) {
	envConfigs := Env{}
	if err := env.Parse(&envConfigs); err != nil {
		log.Fatal(fmt.Errorf("unable to parse configs \n %+v", err))
		return nil, err
	}

	return &envConfigs, nil
}
