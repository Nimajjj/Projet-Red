package main

import (
	"strconv"
	"fmt"
)

var ItemsForSales map[string]int
var ItemList []string

var alreadyDisplayed bool

var MerchantLogs []string

func OpenMerchantMenu() {
	Clear()
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
	if alreadyDisplayed {
		fmt.Print("Merchant: Welcome, how may I help you ?\n")
		fmt.Print(Colorize(Yellow, " 0 - Exit\n"))
	} else {
		SlowPrint("Merchant: Welcome, how may I help you ?\n")
		SlowPrint(Colorize(Yellow, " 0 - Exit\n"))
	}
	i := 0
	for item, price := range ItemsForSales {
		ItemList = append(ItemList, item)
		if alreadyDisplayed {
			fmt.Print(" ", strconv.Itoa(i+1), " - ", item, ": ", strconv.Itoa(price), "$\n")
		} else {
			SlowPrint(" ", strconv.Itoa(i+1), " - ", item, ": ", strconv.Itoa(price), "$\n")
		}
		i++
	}
	for index, log := range MerchantLogs {
		if index == len(MerchantLogs) - 1 {
			SlowPrint(log, "\n")
		} else {
				fmt.Print(log, "\n")
		}
	}
	SlowPrint(Player.Name, "'s money: ", Colorize(Yellow, strconv.Itoa(Player.Money), "$\n"))
	alreadyDisplayed = true
	AskForBuy()
}

func AskForBuy() {
	input := TakeIntInput()
	if input == 0 {
		alreadyDisplayed = false
		MainMenu()
	}
	if OutOfRange(input, 0, len(ItemList)) {
		OpenMerchantMenu()
		return
	}
	MerchantLogs = append(MerchantLogs, Player.buyItem(ItemList[input-1], ItemsForSales[ItemList[input-1]]) )
	OpenMerchantMenu()
}
