package main

import (
	"fmt"
	"log"
)

func main() {
	app := &App{}
	if err := app.Run(); err != nil {
		log.Fatal(fmt.Errorf("unable to run \n %+v", err))
		panic(err)
	}
}
