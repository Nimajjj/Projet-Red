package main

import (
  "log"
  "strconv"
)

func (c Character) displayInventory() {
  invArr := []string{"Exit"}
  SlowPrint(c.Name, "'s inventory:\n"," 0 - Exit\n")
  i := 1
  for item, qt := range c.Inventory {
    SlowPrint(" ", strconv.Itoa(i), " - ", item, ": ", strconv.Itoa(qt), "\n")
    invArr = append(invArr, item)
    i++
  }
  input := TakeIntInput()
  if OutOfRange(input, 0, len(invArr)) {
    c.displayInventory()
    return
  }
  selectItem(invArr[input])
}

func selectItem(item string) {
  switch item {
  case "Exit":
    MainMenu()
  case "Life Potion":
    Player.takePot()
  case "Poison Potion":
    Player.takePoisonPotion()
  case "Spell book: Fire Ball":
    if Player.addSkill("Fire Ball") {
      Player.removeItem("Spell book: Fire Ball")
    }
  }
}

func (c* Character) removeItem(item string) {
  if _, has_item := c.Inventory[item]; has_item {
    c.Inventory[item] -= 1
    if c.Inventory[item] <= 0 {
      delete(c.Inventory, item)
      //fmt.Print(c.Name, " use his last ", item, ".\n")
    }
  } else {
    log.Fatal("Error try to remove unexistant item.")
  }
}

func (c *Character) addItem(item string, quantity int) {
  limit, weight := 10, 0
  for _, qt := range c.Inventory {
    weight += qt
  }
  if weight >= limit {
    SlowPrint("The inventory of ", c.Name, " is full.\n")
    return
  }
  SlowPrint(c.Name, " get ", strconv.Itoa(quantity)," ", item, ".\n")
  c.Inventory[item] += quantity
}

func (c *Character) buyItem(item string) {
  c.addItem(item, 1)
}
