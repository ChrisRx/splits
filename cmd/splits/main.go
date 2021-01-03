package main

import (
	"log"

	"github.com/ChrisRx/splits/cmd/splits/app"
)

func main() {
	if err := app.NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
