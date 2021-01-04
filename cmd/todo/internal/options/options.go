package options

import (
	"github.com/urfave/cli/v2"
)

const (
	markdownFlagName = "markdown"
)

func MarkdownFlag(destination *string, required bool) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        markdownFlagName,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}
