package wire

import "fmt"

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{Name: name}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s kill %s, world peace!\n", m.Player.Name, m.Monster.Name)
}
