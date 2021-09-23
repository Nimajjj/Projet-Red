package main

import (
  "fmt"
  "github.com/TwinProduction/go-color"
  "strconv"
)

var ItemsMap map[string]Item

var LifePotion Item
var PoisonPotion Item


func InitItems() {
  ItemsMap = make(map[string]Item)
  LifePotion.init("Life Potion", "A well-made potion, useful for first aid.\nEffect: +30HP", LifePotionEffect, true)
  PoisonPotion.init("Poison Potion", "A bottle fullish with a dangerous poison.\nEffect: UNKNOWN", PoisonPotionEffect, false)
}


func LifePotionEffect() {
  if Player.Health == Player.HealthMax {
    fmt.Println(Player.Name, "has already all his HP.")
    DisplayMenu(MainMenu)
    return
  }

  fmt.Println(Player.Name, "drink a life potion.")

  Player.Health += 30
  if Player.Health > Player.HealthMax {
    Player.Health = Player.HealthMax
  }

  Player.Inventory["Life Potion"] -= 1
  if Player.Inventory["Life Potion"] == 0 {
    delete(Player.Inventory, "Life Potion")
    fmt.Println(Player.Name, "don't have potions anymore.")
    DisplayMenu(MainMenu)
  }
  fmt.Print(Player.Name, " still have ", Player.Inventory["Life Potion"], " life potion(s).\n")

  fmt.Print(Player.Name, " HP: ")
  health_display := strconv.Itoa(Player.Health) + "/" + strconv.Itoa(Player.HealthMax)
  if Player.Health <= 25 {
    fmt.Println(color.Ize(color.Red, health_display ) )
  } else if Player.Health > 65 {
    fmt.Println(color.Ize(color.Green, health_display ) )
  } else {
    fmt.Println(color.Ize(color.Yellow, health_display ) )
  }
  fmt.Println()
  DisplayMenu(MainMenu)
}



func PoisonPotionEffect() {
  print("you drink the bad potion")
}
