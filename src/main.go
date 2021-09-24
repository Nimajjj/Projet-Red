package main

import (
	"bufio"
	"log"
	"os"
	"flag"
)

var Reader *bufio.Reader
var Player Character

func main() {
	var debugMode = false
	flag.BoolVar(&debugMode, "debug", false, "Running debug mode.")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	Reader = bufio.NewReader(os.Stdin)

	if debugMode {
		InitDefaultCharacter()
	} else {
		NewCharacter()
		WaitEnter()
	}

	MainMenu()
}
