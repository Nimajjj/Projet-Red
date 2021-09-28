package main

import (
  "strconv"
)

var CraftsMaps = map[string]map[string]int{
  "Adventurer Hat": map[string]int {
    "Leather": 2,
    "Crow feathers": 3,
  },
  "Darkest Sword": map[string]int {
    "Wolf fangs": 3,
    "Unicorn horn": 1,
    "Crow feathers": 2,
    "Iron Ingots": 1,
  },
  "Adventurer Cloak": map[string]int {
    "Wolf fur": 2,
    "Leather": 1,
  },
}


func BlacksmithMenu() {
  var craftsList []string
  SlowPrint("Blacksmith: Welcome to my forge!\nI can forge anything for 5$ as long as you have the materials\n")
  SlowPrint(Colorize(Yellow, " 0 - Exit\n\n"))
	i := 0
	for result, craftMap := range CraftsMaps {
    craftsList = append(craftsList, result)

		SlowPrint(" ", strconv.Itoa(i+1), " - ", result, ":\n")

    for ressource, quantity := range craftMap {
      SlowPrint("  └─", ressource, ": ", strconv.Itoa(quantity), "\n")
    }

    println()
		i++
	}
  AskForCraft(craftsList)
}


func (c *Character) craft(item string) {
  if c.Money < 5 {
    SlowPrint(c.Name, " don't have enough money...\n")
    WaitEnter()
    MainMenu()
    return
  }
  for ress, quantity := range CraftsMaps[item] {
    if (c.Inventory[ress] < quantity) {
      SlowPrint(c.Name, " don't have enough materials to craft ", item, "\n")
      WaitEnter()
      return
    }
  }
  for ress, quantity := range CraftsMaps[item] {
    for i := 0; i < quantity; i++ {
      c.removeItem(ress)
    }
  }
  c.addItem(item, 1)
  c.Money -= 5
  SlowPrint(c.Name, " craft a ", item, "\n")
  WaitEnter()
}


func AskForCraft(l []string) {
  input := TakeIntInput()
	if input == 0 {
		MainMenu()
	}
	if OutOfRange(input, 0, len(l)) {
		AskForCraft(l)
		return
	}
  Player.craft(l[input-1])
  AskForCraft(l)
}
