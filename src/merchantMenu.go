package main

import (
	"strconv"
)

var ItemsForSales map[string]int
var ItemList []string

func OpenMerchantMenu() {
	ItemList = []string{}
	ItemsForSales = map[string]int{
		"Life Potion": 14,
		"Iron Sword": 25,
		"Poison Potion": 12,
		"Spell book: Fire Ball": 45,
		"Leather": 3,
		"Iron Ingots": 10,
		"Wolf fur": 10,
		"Wolf fangs": 3,
		"Crow feathers": 1,
		"Unicorn horn": 20,
	}
	SlowPrint("Merchant: Welcome, how may I help you ?\n")
	SlowPrint(Colorize(Yellow, " 0 - Exit\n"))
	i := 0
	for item, price := range ItemsForSales {
		ItemList = append(ItemList, item)
		SlowPrint(" ", strconv.Itoa(i+1), " - ", item, ": ", strconv.Itoa(price), "$\n")
		i++
	}
	AskForBuy()
}

func AskForBuy() {
	input := TakeIntInput()
	if input == 0 {
		MainMenu()
	}
	if OutOfRange(input, 0, len(ItemList)) {
		OpenMerchantMenu()
		return
	}
	Player.buyItem(ItemList[input-1], ItemsForSales[ItemList[input-1]])
	AskForBuy()
}
