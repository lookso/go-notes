package main

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "sex",
				Aliases: []string{"s"},
				Usage:   "你可以通过我来获取性别",
				Action: func(c *cli.Context) error {
					fmt.Println("难人")
					return nil
				},
			},
			{
				Name:    "age",
				Aliases: []string{"a"},
				Usage:   "你可以通过我来获取年龄",
				Action: func(c *cli.Context) error {
					fmt.Println("100岁了")
					return nil
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands)) // 通过命令函数来排序，在help中进行展示

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
