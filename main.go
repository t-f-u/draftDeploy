package main

import (
	"draftDeploy/app"
	"errors"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal(errors.New("there must be 2 arguments: target and distribution"))
	}

	if err := app.Start(os.Args[1], os.Args[2]); err != nil {
		log.Fatal(err.Error())
	}
}
