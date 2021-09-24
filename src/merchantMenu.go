package main

import (
  "strconv"
)

func OpenMerchantMenu() {
  itemsForSales := []string{"Life Potion", "Iron Sword", "Poison Potion", "Spell book: Fire Ball"}
  SlowPrint("Merchant: Welcome, how may I help you ?\n")
  SlowPrint(" 0 - Exit\n")
  i := 0
  for _, item := range itemsForSales {
    SlowPrint(" ", strconv.Itoa(i+1), " - ", item, "\n")
    i++
  }
  input := TakeIntInput()
  if input == 0 {
    MainMenu()
  }
  if OutOfRange(input, 0, len(itemsForSales)) {
    OpenMerchantMenu()
    return
  }
  Player.buyItem(itemsForSales[input-1])
  MainMenu()
}
