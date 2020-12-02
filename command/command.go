package command

import (
	"github.com/urfave/cli/v2"
)

var registeredCommands = make(cli.Commands, 0)
var registeredSubCommands = map[string]cli.Commands{}

func registerCommand(cmd *cli.Command) {
	registeredCommands = append(registeredCommands, cmd)
}

func registerSubCommand(parent string, cmd *cli.Command) {
	if registeredSubCommands[parent] == nil {
		registeredSubCommands[parent] = make(cli.Commands, 0)
	}

	registeredSubCommands[parent] = append(registeredSubCommands[parent], cmd)
}

// Metadata constants set in the CLI app.
const (
	BinaryName     = "response"
	AppDescription = "A performant, feature-packed, and extensible MDT, CAD, and RMS for roleplay communtiies."
	AppUsage       = "Start, manage and interact with Response."
)

// New initalizes and creates a new cli.App instance that can be used
// to the
func New() *cli.App {
	app := cli.NewApp()

	// set the metadata
	app.Name = BinaryName
	app.Description = AppDescription
	app.Usage = AppUsage

	// set options
	app.EnableBashCompletion = true

	// global flags
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "use a specific configuration file",
			DefaultText: "",
		},
	}

	// initialize things that matter, like our configuration!
	app.Before = func(ctx *cli.Context) error {
		// if err := config.InitWithPath(ctx.String("config")); err != nil {
		// 	return err
		// }

		// return nil
		return nil
	}

	// load the registered commands
	app.Commands = registeredCommands
	for _, cmd := range app.Commands {
		if registeredSubCommands[cmd.Name] != nil {
			cmd.Subcommands = registeredSubCommands[cmd.Name]
		}
	}

	return app
}
