package game

import (
	"os/exec"
)

type Game struct{}

func Start() {
	exec.Command("clear")
}
