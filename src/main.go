package main

import (
	"bufio"
	"os"
)

var Reader *bufio.Reader
var Player Character

var Debug bool
var Slow bool
var BootState string
var Intro bool

func main() {
	DebugInit()

	Reader = bufio.NewReader(os.Stdin)

	InitEquipments()
	InitEnemies()

	if Debug || Intro {	// -d: Init debug mode -> default character & no slow print
		InitDefaultCharacter()
	} else {
		NewCharacter()
		WaitEnter()
	}

	switch BootState {	// -boot=case
	case "f":
		InitFight([]Enemy{Gnom, Gnom}, "Training Field")
	case "m":
		OpenMerchantMenu()
	case "i":
		Player.displayInventory(false)
	case "b":
		BlacksmithMenu()
	}

	MainMenu()
}
