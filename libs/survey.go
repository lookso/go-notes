package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Hello, World!")
	var isUpDb string
	prompt := &survey.Input{
		Message: "本次运行是否需要更新覆盖数据? (输入 '1' 更新，回车不更新):",
	}

	// 提示用户输入
	err := survey.AskOne(prompt, &isUpDb)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if isUpDb == "1" {
		fmt.Println("更新覆盖数据")
		// 执行更新覆盖数据的逻辑
	} else {
		fmt.Println("不更新覆盖数据")
		// 跳过更新覆盖数据的逻辑
	}

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

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
