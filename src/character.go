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
  Skills []string
}


func (c *Character) init(name, race string, level, max_health, health int, inventory map[string]int, skills []string){
  c.Name = name
  c.Race = race
  c.Level = level
  c.HealthMax = max_health
  c.Health = health
  c.Inventory = inventory
  c.Skills = skills
}


func (c Character) displaySheet() {
  SlowPrint(c.Name, "'s sheet:\n")
  SlowPrint(" Name: ", c.Name, "\n")
  SlowPrint(" Race: ", c.Race, "\n")
  SlowPrint(" Health: ")
  c.printHealth()
  SlowPrint(" Skills:\n")
  for _, skill := range(c.Skills) {
    SlowPrint("   -", skill, "\n")
  }
  WaitEnter()
}


func (c *Character) takePot() {
  if c.Health == c.HealthMax {
    SlowPrint(c.Name, " has already all his HP.\n")
    return
  }

  SlowPrint(c.Name, " drink a life potion.\n")

  c.addToHealth(30)

  SlowPrint(c.Name, "'s HP: ")
  c.printHealth()

  if c.Inventory["Life Potion"] != 0 {
    c.removeItem("Life Potion")
  } else {
    SlowPrint(c.Name, " still have ", strconv.Itoa(c.Inventory["Life Potion"]), " life potion(s).\n")
  }

  fmt.Println()
}

func (c Character) printHealth() {
  health_display := strconv.Itoa(c.Health) + "/" + strconv.Itoa(c.HealthMax) + "\n"
  if c.Health <= (c.HealthMax / 3) {
    SlowPrint(color.Ize(color.Red, health_display ) )
  } else if c.Health > (c.HealthMax / 3) * 2 {
    SlowPrint(color.Ize(color.Green, health_display ) )
  } else {
    SlowPrint(color.Ize(color.Yellow, health_display ) )
  }
}

func (c *Character) dead() {
  if c.Health <= 0 {
    str := c.Name
    str += " is dead ...\n"
    SlowPrint(color.Ize(color.Red, str ) )
    c.Health = c.HealthMax / 2
    SlowPrint("Fortunately, a priest who passed by was able to resurrect you.\nDon't waste your second chance.\n")
  }
}

func (c *Character) addToHealth(qt int) {
  c.Health += qt

  if c.Health > c.HealthMax {
    c.Health = c.HealthMax
  }

  c.dead()
}


func (c *Character) addSkill(skill string) bool {
  if !IsStrInArray(skill, c.Skills) {
    c.Skills = append(c.Skills, skill)
    SlowPrint(c.Name, " learned the skill: ", skill, "\n")
    return true
  } else {
    SlowPrint(c.Name, " already knows the skill: ", skill, "\n")
    return false
  }
}


func InitDefaultCharacter() {
  Player.init("Benjamin", "Human", 1, 100, 40, map[string]int{"Life Potion":3, "Poison Potion":1}, []string{"Punch"})
}
