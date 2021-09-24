package main

import (
  "time"
  "fmt"
)

func (c *Character) takePoisonPotion() {
  for i := 0; i != 3; i++ {
    fmt.Print(c.Name, " take 10dmg due to poison.\n")
    c.addToHealth(-10)
    c.printHealth()
    time.Sleep(1 * time.Second)
  }
}
