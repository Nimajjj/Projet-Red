package main

import (
	"time"
)

func (c *Character) takePoisonPotion() {
	dead := false
	for i := 0; i != 3; i++ {
		SlowPrint(c.Name, " take 10dmg due to poison.\n")
		dead = c.addToHealth(-10)
		if dead {
			break
		}
		c.printHealth()
		time.Sleep(1 * time.Second)
	}
	/*if !dead {
		WaitEnter()
	}*/
}
