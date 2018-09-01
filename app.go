package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/jianhan/gopt/handler"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

type config struct {
	Addr         string `env:"ADDR"`
	WriteTimeout int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	ReadTimeout  int    `env:"READ_TIMEOUT" envDefault:"15"`
	IdleTimeout  int    `env:"IDLE_TIMEOUT" envDefault:"15"`
}

type App struct {
}

func (a *App) Run() error {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return err
	}

	// get configs
	cfg := config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to parse configs \n %+v", err))
		return err
	}

	// get router from handler
	r, err := handler.NewRouter()
	if err != nil {
		log.Fatal(fmt.Errorf("unable to init router %v", err))
		return err
	}

	// get router
	router, err := r.GetRouter()
	if err != nil {
		log.Fatal(fmt.Errorf("unable to get router %v", err))
		return err
	}

	// init server
	srv := &http.Server{
		Handler: router,
		Addr:    cfg.Addr,
		// settings
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	return nil
}

// init handles app configs ,etc... before boot
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		panic(err)
	}
}
