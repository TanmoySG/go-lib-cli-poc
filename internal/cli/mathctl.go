package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    var language string

    app := &cli.App{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:        "add",
                Value:       "english",
                Usage:       "to do",
                Destination: &language,
            },
        },
        Action: func(cCtx *cli.Context) error {
            name := "someone"
            if cCtx.NArg() > 0 {
                name = cCtx.Args().Get(0)
            }
            if language == "spanish" {
                fmt.Println("Hola", name)
            } else {
                fmt.Println("Hello", name)
            }
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
