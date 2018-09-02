package main

import (
	"fmt"
	"github.com/jianhan/gopt/config"
	"github.com/jianhan/gopt/handler"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"time"
)

type App struct {
}

func (a *App) Run() error {
	// load env
	envConfigs, err := config.EnvConfigs()
	if err != nil {
		panic(err)
	}

	// get router from handler
	// define middleware pass into it
	r, err := handler.NewRouter(
		[]negroni.Handler{negroni.NewRecovery(), negroni.NewLogger()},
		[]handler.APIRouter{handler.NewUser()},
	)
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
		Handler: cors.Default().Handler(router),
		Addr:    envConfigs.Address(),
		// settings
		WriteTimeout: time.Duration(envConfigs.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(envConfigs.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(envConfigs.IdleTimeout) * time.Second,
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
