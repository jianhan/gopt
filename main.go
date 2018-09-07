package main

import (
	"fmt"
	"github.com/jianhan/gopt/place"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	client, err := place.GetClient()
	if err != nil {
		panic(err)
	}
	logrus.Info(client)
	return

	app := &App{}
	if err := app.Run(); err != nil {
		log.Fatal(fmt.Errorf("unable to run \n %+v", err))
		panic(err)
	}
}
