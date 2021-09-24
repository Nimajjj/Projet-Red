package main

import (
	"bufio"
	"os"
)

var Reader *bufio.Reader
var Player Character
var Debug bool

func main() {
	DebugInit()

	Reader = bufio.NewReader(os.Stdin)

	if Debug {
		InitDefaultCharacter()
	} else {
		NewCharacter()
		WaitEnter()
	}

	MainMenu()
}
