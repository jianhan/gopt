package main

import (
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/jianhan/gopt/middleware"
	"github.com/jianhan/gopt/place"
	"github.com/sirupsen/logrus"
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
	googleClient, err := place.GetClient()
	if err != nil {
		panic(err)
	}

	// create a limiter struct.
	limiter := tollbooth.NewLimiter(10000, nil)

	// load env
	envConfigs, err := config.EnvConfigs()
	if err != nil {
		panic(err)
	}

	// setup cache
	cache, cErr := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if cErr != nil {
		panic(cErr)
	}

	// get router from handler
	// define middleware pass into it
	r, err := handler.NewRouter(
		[]negroni.Handler{
			negroni.NewRecovery(),
			negroni.NewLogger(),
			negroni.HandlerFunc(middleware.JSONRspHeader),
			tollbooth_negroni.LimitHandler(limiter),
		},
		[]handler.APIRouter{
			handler.NewUser(),
			handler.NewPlace(googleClient, cache),
		},
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
	var debug bool
	if envConfigs.Environment == "development" {
		debug = true
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8080"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
		Debug:            debug,
	})
	logrus.Info(c)
	srv := &http.Server{
		Handler: router,
		Addr:    envConfigs.Address(),
		// settings
		WriteTimeout: time.Duration(envConfigs.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(envConfigs.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(envConfigs.IdleTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	return nil
}
