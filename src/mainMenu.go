package main

import (
  "os"
  "strconv"
)

func MainMenu() {
  os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
  SlowPrint("Main Menu :\n")
  Options := []string{"Character Sheet", "Inventory", "Merchant", "Exit"}
  for i, o := range Options {
    SlowPrint(" ", strconv.Itoa(i), " - ", o, "\n")
  }
  i := TakeIntInput()

  if !OutOfRange(i, 0, len(Options)-1) {
    os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
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
