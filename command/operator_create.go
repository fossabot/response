package command

import (
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jedib0t/go-pretty/text"
	"github.com/responserms/response/operator"
	"github.com/urfave/cli/v2"
)

type operatorCreateAnswers struct {
	Name            string
	Path            string
	EnableDeveloper bool
	EncryptionKey   string
}

func handleOperatorCreateCommand(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		lineColor(text.FgHiRed, "\nYou need to specify the path for Response to be initialized.")
		lineColor(text.FgRed, "example: response operator create my-response")
	}

	op, err := operator.New()
	if err != nil {
		lineColor(text.FgHiRed, "\n We could not initialize the Response operator. Error: %s", err.Error())
	}

	path := path.Join("./", ctx.Args().First())
	if p := ctx.String("path"); p != "" {
		path = p
	}

	opAnswers := &operatorCreateAnswers{
		Name: ctx.Args().Get(0),
		Path: path,
	}

	line("\nCreating a Response instance in %q with the name of %q.\n", path, ctx.Args().First())

	survey.AskOne(
		&survey.Confirm{
			Message: "Enable global developer mode?",
			Help:    "If enabled, this will enable additional logging. When enabled globally like this, every user will see the UI developer tools. If disabled, specific users that have the necessary permissions can enable developer mode for themselves.",
			Default: false,
		},
		&opAnswers.EnableDeveloper,
	)

	var provideEncryptionKey = false
	survey.AskOne(
		&survey.Confirm{
			Message: "Would you like to provide your own encryption key?",
			Help:    "This is useful if migrating Response. Using a different encryption key will make data that has been encrypted before inaccessible.",
			Default: false,
		},
		&provideEncryptionKey,
	)

	if provideEncryptionKey {
		survey.AskOne(
			&survey.Input{
				Message: "Provide your encryption key",
				Help:    "You opted to provide an encryption key, please provide it here. Otherwise cancel and restart but choose to use a generated key.",
				Default: "",
			},
			&opAnswers.EncryptionKey,
		)

		if len(opAnswers.EncryptionKey) < 32 {
			lineColor(text.FgHiYellow, "\nWARN: The encryption key you provided is less than 32 characters. This is potentially unsafe.\n")
		}
	} else {
		opAnswers.EncryptionKey = operator.GenEncryptionKey()
	}

	op.CreateLocal(ctx.Args().First(), &operator.CreateLocalOptions{
		Path:            path,
		EnableDeveloper: opAnswers.EnableDeveloper,
		EncryptionKey:   opAnswers.EncryptionKey,
	})

	return nil
}

func init() {
	registerSubCommand("operator", &cli.Command{
		Name:        "create",
		Description: "Create a configuration file and data_dir for Response in the given directory.",
		Action:      handleOperatorCreateCommand,
		Usage:       "Initialize a new Response instance with the given name",
		ArgsUsage:   "<name>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Usage:       "path to initialize Response",
				DefaultText: "./<name>",
			},
		},
	})
}
