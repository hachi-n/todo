package main

import (
	"github.com/hachi-n/todo/cmd/todo/internal/options"
	entrypoints "github.com/hachi-n/todo/lib/entrypoints/register"
	"github.com/urfave/cli/v2"
)

func registerCommand() *cli.Command {
	var markdown string
	return &cli.Command{
		Name:        "register",
		Usage:       "todo register",
		Description: "todo register.",
		Flags: []cli.Flag{
			options.MarkdownFlag(&markdown, false),
		},
		Action: func(c *cli.Context) error {
			return entrypoints.Apply(markdown)
		},
	}
}
