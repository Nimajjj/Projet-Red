package main

import (
  "fmt"
  "os"
)

var MainMenu Menu

type Choice struct {
  Name string
  FLink func()  // activated func on use
  FLinkArg func(interface{})  // alternate func usable but with an arg
  Args [](interface{})        // arg for the FLinkArg func
}

type Menu struct {
  Name string
  ChoiceArr []Choice
}


func InitMenus() {
  characterInfo := Choice{"Character Sheet", CharacterSheet, NullArg, NullAI}
  inventory := Choice{"Inventory", Player.displayInventory, NullArg, NullAI}
  merchant := Choice{"Merchant", OpenMerchantMenu, NullArg, NullAI}
  quit := Choice{"Quit", Quit, NullArg, NullAI}

  choiceMainMenu := []Choice{characterInfo, inventory, merchant, quit}
  MainMenu = Menu{"MainMenu", choiceMainMenu}
}


func DisplayMenu(MenuDisplayed Menu) {
  fmt.Print(MenuDisplayed.Name, " :\n")
  for i, choice := range MenuDisplayed.ChoiceArr {
    fmt.Print(i, " - ", choice.Name, "\n")
  }
  ChoiceHandler(MenuDisplayed)
}


func ChoiceHandler(MenuDisplayed Menu) {
  choice := TakeIntInput()
  if choice == -1 || choice > len(MenuDisplayed.ChoiceArr) - 1 {
    fmt.Println("Bad Input... Type 'exit' if stuck in loop.\nOr press 'enter' to come back to the last menu.")
    if TakeStrInput() == "exit" {
      os.Exit(0)
    }
    DisplayMenu(MenuDisplayed)
    return
  }
  if len(MenuDisplayed.ChoiceArr[choice].Args) > 0 {
    arg := MenuDisplayed.ChoiceArr[choice].Args[0]
    MenuDisplayed.ChoiceArr[choice].FLinkArg(arg)
  } else {
    MenuDisplayed.ChoiceArr[choice].FLink()
  }
}


func CharacterSheet() {
  Player.displaySheet()
  DisplayMenu(MainMenu)
}

func Quit() {
  os.Exit(0)
}
