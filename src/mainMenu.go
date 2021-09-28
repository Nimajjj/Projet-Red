package main

import (
	"os"
	"strconv"
)

func MainMenu() {
	Clear()
	SlowPrint("Main Menu :\n")
	Options := []string{"Character Sheet", "Inventory", "Merchant", "Blacksmith", "Training Field", "Exit"}
	for i, o := range Options {
		if o == "Exit" {
			SlowPrint(Colorize(Yellow, " ", strconv.Itoa(i), " - ", o, "\n"))
			continue
		}
		SlowPrint(" ", strconv.Itoa(i), " - ", o, "\n")
	}

	i := TakeIntInput()

	if !OutOfRange(i, 0, len(Options)-1) {
		Clear()
		switch Options[i] {
		case "Character Sheet":
			Player.displaySheet()
		case "Inventory":
			Player.displayInventory()
		case "Merchant":
			OpenMerchantMenu()
		case "Blacksmith":
			BlacksmithMenu()
		case "Training Field":
			InitFight([]Enemy{Goblin, Gnom, Gnom, Goblin}, "Training Field")
		case "Exit":
			os.Exit(0)
		}
	}
	MainMenu()
}
