package main

import (
	"log"
	"strconv"
)

func (c Character) displayInventory() {
	invArr := []string{"Exit"}
	SlowPrint(c.Name, "'s inventory:\n", Colorize(Yellow, " 0 - Exit\n"))
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
	default:
		if _, isIn := EquipmentList[item]; isIn {
			Player.equip(EquipmentList[item].Slot, item)
		}
	}
	WaitEnter()
}

func (c *Character) equip(slot, item string) {
	replace := false
	if c.hasEquipment(slot) {
		c.addItem(c.Equipment[slot].Name, 1)
		c.HealthMax -= c.Equipment[slot].Effect
		replace = true
	}

	if replace {
		SlowPrint(c.Name, " replace his '", c.Equipment[slot].Name, "' with '", item, "'.\n")
	} else {
		SlowPrint(c.Name, " equipped '", item, "'.\n")
	}

	c.Equipment[slot] = EquipmentList[item]
	c.HealthMax += c.Equipment[slot].Effect
	c.addToHealth(0)
	c.removeItem(item)
}

func (c Character) hasEquipment(slot string) bool {
	if c.Equipment[slot] != Empty{
		return true
	}
	return false
}

func (c *Character) removeItem(item string) {
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

func (c *Character) addItem(item string, quantity int) bool {
	limit, weight := 100, 0
	for _, qt := range c.Inventory {
		weight += qt
	}
	if weight >= limit {
		SlowPrint(Colorize(Red, "The inventory of ", c.Name, " is full.\n"))
		return false
	}
	c.Inventory[item] += quantity
	return true
}

func (c *Character) buyItem(item string, price int) {
	if c.Money - price <= 0 {
		SlowPrint(c.Name, " don't have enough money to buy '", item, "'.\n")
		return
	}
	if c.addItem(item, 1) {
		c.Money -= price
		SlowPrint(c.Name, " buy one '", item, "'.\n")
	}
}
