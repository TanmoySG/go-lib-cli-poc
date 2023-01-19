package mathctl

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/TanmoySG/go-lib-cli-poc/internal/mathctlFuncs"
)

func main() {

	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	addFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "n1",
			Value:       "english",
			Usage:       "to do",
		},
		&cli.StringFlag{
			Name:        "n2",
			Value:       "english",
			Usage:       "to do",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "add",
			Flags: addFlags,
			Action: mathctlFuncs.Add,
		},
		{
			Name: "sub",
			Flags: addFlags,
			Action: mathctlFuncs.Sub,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
