package main

import (
	"log"

	"github.com/NonsoAmadi10/lightning-web-app/app"
)

func main() {
	err := app.App().Start("localhost:433")
	if err != nil {
		log.Fatal(err)
	}
}
