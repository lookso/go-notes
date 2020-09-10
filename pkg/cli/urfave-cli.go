package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "GoTool"
	//程序的用途描述
	oApp.Usage = "To save the world"
	//程序的版本号
	oApp.Version = "1.0.0"

	//预置变量
	var host string
	var debug bool
	var env string

	//设置启动参数
	oApp.Flags = []cli.Flag{
		//参数类型string,int,bool
		cli.StringFlag{
			Name:        "host",           //参数名字
			Value:       "127.0.0.1",      //参数默认值
			Usage:       "Server Address", //参数功能描述
			Destination: &host,            //接收值的变量
		},
		cli.IntFlag{
			Name:  "port,p",
			Value: 8888,
			Usage: "Server port",
		},
		cli.StringFlag{
			Name:        "env,e",
			Value:       "prod",
			Usage:       "run env",
			Destination: &env,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug mode",
			Destination: &debug,
		},
	}
	//该程序执行的代码
	oApp.Action = func(c *cli.Context) error {
		fmt.Printf("host=%v \n", host)
		fmt.Printf("host=%v \n", c.Int("port")) //不使用变量接收，直接解析
		fmt.Printf("host=%v \n", debug)
		fmt.Printf("env=%v \n", env)
		return nil
	}
	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	// go run ./urfave-cli.go --env dev
}
