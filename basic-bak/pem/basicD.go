package main

import "fmt"

type Cmd struct {
	Ps    string
	Ls    string
	Pwd   string
	State bool
}

func NewCmd() *Cmd {
	client := &Cmd{}
	return client
}
func (c *Cmd) Init() {
	c.State = true
	c.Ls = "LS"
	c.Ps = "PS"
	c.Pwd = "PWD"
}
func (c *Cmd) Run() {
	c.Init()
}
func main() {
	var cmd = NewCmd()
	cmd.Run()
	fmt.Println(cmd.Pwd)
}
