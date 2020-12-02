package command

import (
	"github.com/jedib0t/go-pretty/text"
	"github.com/responserms/response/operator"
	"github.com/urfave/cli/v2"
)

type operatorInitAnswers struct{}

var operatorInitMessage = `
Initializing the Response operator. The operator...
  - allows you to manage many independent Response instances (like dev, test, and prod)
  - manages locally installed instances (and cloud instances in the future)
  - allows importing (and exporing) datasets in Response
  - allows managing database backups and restores
  and more...`

func handleOperatorInitCommand(ctx *cli.Context) error {
	line(operatorInitMessage)

	o, err := operator.New()
	if err != nil {
		lineColor(text.FgHiRed, "Whoops! There was a problem with initializing the operator: %s", err.Error())
		return err
	}
	defer o.Close()

	lineColor(text.FgHiGreen, "\nThe operator was initialized.")

	return nil
}

func init() {
	registerSubCommand("operator", &cli.Command{
		Name:        "init",
		Description: "Initialize the Response operator on this machine for this user",
		Action:      handleOperatorInitCommand,
		Usage:       "Initialize the Response operator on this machine for this user",
	})
}
