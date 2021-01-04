package main

import (
	entrypoints "github.com/hachi-n/todo/lib/entrypoints/list"
	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Usage:       "todo list",
		Description: "todo list.",
		Flags: []cli.Flag{
		},
		Action: func(c *cli.Context) error {
			return entrypoints.Apply()
		},
	}
}
