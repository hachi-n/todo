package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// github

func main()  {
	app := &cli.App{
		Name: "todo",
		Usage: "todo [sub commands] [flags]",
		Description: "",
		Commands: []*cli.Command{
			registerCommand(),
			listCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
