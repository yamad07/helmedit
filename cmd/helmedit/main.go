package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "helmedit",
		Usage:  "edit helm",
		Action: edit,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
