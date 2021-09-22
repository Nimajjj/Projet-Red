package main

import (
  "fmt"
  "os"
)

var MainMenu Menu

type Choice struct {
  Name string
  FLink func()
}

type Menu struct {
  Name string
  ChoiceArr []Choice
}


func InitMenus() {
  CharacterInfo := Choice{"Character Sheet", CharacterSheet}
  Inventory := Choice{"Inventory", Inventory}
  Quit := Choice{"Quit", Quit}

  ChoiceMainMenu := []Choice{CharacterInfo, Inventory, Quit}
  MainMenu = Menu{"MainMenu", ChoiceMainMenu}
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
    fmt.Println("Bad Input... Type 'exit' if stuck in loop.\n")
    DisplayMenu(MenuDisplayed)
    return
  }
  MenuDisplayed.ChoiceArr[choice].FLink()
}


func CharacterSheet() {
  Player.displaySheet()
  WaitEnter()
}


func Inventory() {
  println("Inventory")
}


func Quit() {
  os.Exit(0)
}
