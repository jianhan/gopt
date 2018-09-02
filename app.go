package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_negroni"
	"github.com/jianhan/gopt/config"
	"github.com/jianhan/gopt/handler"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

type App struct {
}

func (a *App) Run() error {
	// create a limiter struct.
	limiter := tollbooth.NewLimiter(10000, nil)

	// load env
	envConfigs, err := config.EnvConfigs()
	if err != nil {
		panic(err)
	}
	// Create a limiter struct.

	// get router from handler
	// define middleware pass into it
	r, err := handler.NewRouter(
		[]negroni.Handler{
			negroni.NewRecovery(),
			negroni.NewLogger(),
			tollbooth_negroni.LimitHandler(limiter),
		},
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
