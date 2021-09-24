package main

import (
	"strconv"
)

var ItemsForSales []string

func OpenMerchantMenu() {
	ItemsForSales = []string{"Life Potion", "Iron Sword", "Poison Potion", "Spell book: Fire Ball"}
	SlowPrint("Merchant: Welcome, how may I help you ?\n")
	SlowPrint(Colorize(Yellow, " 0 - Exit\n"))
	i := 0
	for _, item := range ItemsForSales {
		SlowPrint(" ", strconv.Itoa(i+1), " - ", item, "\n")
		i++
	}
	AskForBuy()
}

func AskForBuy() {
	input := TakeIntInput()
	if input == 0 {
		MainMenu()
	}
	if OutOfRange(input, 0, len(ItemsForSales)) {
		OpenMerchantMenu()
		return
	}
	Player.buyItem(ItemsForSales[input-1])
	AskForBuy()
}
