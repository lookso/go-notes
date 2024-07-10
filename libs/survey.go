package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "interactive-cli",
		Usage: "A simple CLI with interactive prompt",
		Commands: []*cli.Command{
			{
				Name:  "ask",
				Usage: "Ask a confirm question",
				Action: func(c *cli.Context) error {
					override := false
					prompt := &survey.Confirm{
						Message: "Do you want to override the folder?",
					}
					err := survey.AskOne(prompt, &override)
					if err != nil {
						return err
					}
					fmt.Printf("Override: %t\n", override)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
