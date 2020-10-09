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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "values",
				Aliases: []string{"v"},
				Value:   "",
				Usage:   "set key-values like image.tag.repository=nginx,tag=latest",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "values.yaml",
				Usage:   "set file",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
