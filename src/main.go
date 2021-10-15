package main

/*
	Projet-Red 2021 Ynov
	Borello Benjamin

	Use Powershell 7 for better compatibility (color && cls)
	'go run . --help' for all debug flags
	'go run . -s' to disable slow print (you certainly saw it enough...)
*/


import (
	"bufio"
	"os"
	"syscall"
)

var Reader *bufio.Reader
var Player Character

var Debug bool
var Slow bool
var BootState string
var Intro bool

func main() {
	handle := syscall.Handle(os.Stdout.Fd())	// fucked up
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004)
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
		InitFight([]Enemy{Gnom, Goblin}, "Training Field")
	case "m":
		OpenMerchantMenu()
	case "i":
		Player.displayInventory(false)
	case "b":
		BlacksmithMenu()
	}

	MainMenu()
}
