package main

import (
	"log"

	"github.com/mclark4386/personal-site/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
