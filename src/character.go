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
  fmt.Println("Name: ", c.Name)
  fmt.Println("Race: ", c.Race)
  fmt.Print("Health: ")
  c.printHealth()
  fmt.Print("Skills:\n")
  for _, skill := range(c.Skills) {
    fmt.Print(" -", skill, "\n")
  }
  fmt.Println()
}


func (c *Character) takePot() {
  if c.Health == c.HealthMax {
    fmt.Println(c.Name, "has already all his HP.")
    return
  }

  fmt.Println(c.Name, "drink a life potion.")

  c.addToHealth(30)

  fmt.Print(c.Name, "'s HP: ")
  c.printHealth()

  if c.Inventory["Life Potion"] != 0 {
    c.removeItem("Life Potion")
  } else {
    fmt.Print(c.Name, " still have ", c.Inventory["Life Potion"], " life potion(s).\n")
  }

  fmt.Println()
}

func (c Character) printHealth() {
  health_display := strconv.Itoa(c.Health) + "/" + strconv.Itoa(c.HealthMax)
  if c.Health <= 25 {
    fmt.Println(color.Ize(color.Red, health_display ) )
  } else if c.Health > 65 {
    fmt.Println(color.Ize(color.Green, health_display ) )
  } else {
    fmt.Println(color.Ize(color.Yellow, health_display ) )
  }
}

func (c *Character) dead() {
  if c.Health <= 0 {
    str := c.Name
    str += " is dead ...\n"
    fmt.Print(color.Ize(color.Red, str ) )
    c.Health = c.HealthMax / 2
    fmt.Print("Fortunately, a priest who passed by was able to resurrect you.\nDon't waste your second chance.\n")
  }
}

func (c *Character) addToHealth(qt int) {
  c.Health += qt

  if c.Health > c.HealthMax {
    c.Health = c.HealthMax
  }

  c.dead()
}


func (c *Character) addSkill(skill string) {
  if !IsStrInArray(skill, c.Skills) {
    c.Skills = append(c.Skills, skill)
    fmt.Print(c.Name, " learned the skill: ", skill, "\n")
  } else {
    fmt.Print(c.Name, " already knows the skill: ", skill, "\n")
  }
}


func InitCharacter() {
  Player.init("Benjamin", "Elfe", 1, 100, 40, map[string]int{"Life Potion":3, "Poison Potion":1}, []string{"Punch"})
}
