package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/TanmoySG/go-lib-cli-poc/internal/mathctlFuncs"
)

func main() {

	var n1,n2 string

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "n1",
			Value: "english",
			Usage: "to do",
			Destination: &n1,
		},
		&cli.StringFlag{
			Name:  "n2",
			Value: "english",
			Usage: "to do",
			Destination: &n2,
		},
	},

	app := &cli.App{
		Commands: cli.Commands{
			Flags: flags,

		}
		Action: mathctlFuncs.Add,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
