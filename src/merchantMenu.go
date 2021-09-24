package main

import (
  "fmt"
)

func OpenMerchantMenu() {
  itemsForSales := []string{"Life Potion", "Iron Sword", "Poison Potion", "Spell book: Fire Ball"}
  fmt.Print("Merchant: Welcome, how may I help you ?\n")
  fmt.Print(" 0 - Exit\n")
  i := 0
  for _, item := range itemsForSales {
    fmt.Print(" ", i+1, " - ", item, "\n")
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
