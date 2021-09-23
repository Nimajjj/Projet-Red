package main

import (
  "fmt"
  "github.com/TwinProduction/go-color"
  "strconv"
)

type Character struct {
  Name string
  Race string
  Level int
  HealthMax int
  Health int
  Inventory map[string]int
}


func (c *Character) init(name, race string, level, max_health, health int, inventory map[string]int){
  c.Name = name
  c.Race = race
  c.Level = level
  c.HealthMax = max_health
  c.Health = health
  c.Inventory = inventory
}


func (c Character) displaySheet() {
  fmt.Println("Name: ", c.Name)
  fmt.Println("Race: ", c.Race)
  health_display := strconv.Itoa(c.Health) + "/" + strconv.Itoa(c.HealthMax)
  fmt.Print("Health: ")
  if c.Health <= 25 {
    fmt.Println(color.Ize(color.Red, health_display ) )
  } else if c.Health > 65 {
    fmt.Println(color.Ize(color.Green, health_display ) )
  } else {
    fmt.Println(color.Ize(color.Yellow, health_display ) )
  }
  fmt.Println()
}

func (c Character) displayInventory() {
  var choiceMenuArr []Choice
  var exitChoice = Choice{color.Ize(color.Yellow, "Exit"), DisplayMainMenu, NullArg, []interface{}{}}

  choiceMenuArr = make([]Choice, len(c.Inventory) - 2)
  choiceMenuArr = append(choiceMenuArr, exitChoice)

  for item_name := range c.Inventory {
    var arg = []interface{}{}
    arg = append(arg, item_name)
    var itemChoice = Choice{item_name, Null, SelectItem, arg}
    choiceMenuArr = append(choiceMenuArr, itemChoice)
  }

  var Inventory = Menu{"Inventory", choiceMenuArr}
  DisplayMenu(Inventory)
}


func InitCharacter() {
  Player.init("Benjamin", "Elfe", 1, 100, 40, map[string]int{"Life Potion":3, "Poison Potion":1})
}
