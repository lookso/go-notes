package wire

import "testing"

func TestOldMission(t *testing.T) {
	monster := NewMonster("Smaug")
	player := NewPlayer("Bilbo Baggins")
	mission := NewMission(player, monster)

	mission.Start()
}

func TestNewMission(t *testing.T) {
	mission := InitMission("Bilbo Baggins", "Smaug")
	mission.Start()
}
