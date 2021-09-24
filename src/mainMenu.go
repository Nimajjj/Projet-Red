package main

import (
  "fmt"
  "os"
)

func MainMenu() {
  fmt.Print("Main Menu :\n")
  Options := []string{"Character Sheet", "Inventory", "Merchant", "Exit"}
  for i, o := range Options {
    fmt.Print(i, " - ", o, "\n")
  }
  i := TakeIntInput()

  if !OutOfRange(i, 0, len(Options)-1) {
    switch Options[i] {
    case "Character Sheet":
      Player.displaySheet()
    case "Inventory":
      Player.displayInventory()
    case "Merchant":
      OpenMerchantMenu()
    case "Exit":
      os.Exit(0)
    }
  }
  MainMenu()
}
