package command

import (
	"github.com/urfave/cli/v2"
)

func init() {
	registerCommand(&cli.Command{
		Name:        "operator",
		Description: "Operate a Response instance installed on this machine.",
		Flags:       []cli.Flag{},
		Usage:       "Manage a Response instance on this machine.",
	})
}
