package main

import (
  "fmt"
)

type Character struct {
  Name string
  Race string
  Level int
  HealthMax int
  Health int
  Inventory map[string]int
}


func (c *Character) Init(name, race string, level, max_health, health int, inventory map[string]int){
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
  fmt.Println("Health: ", c.Health, "/", c.HealthMax)
}

func InitCharacter() {
  Player.Init("Benjamin", "Elfe", 1, 100, 40, map[string]int{"Potions": 3})
}
