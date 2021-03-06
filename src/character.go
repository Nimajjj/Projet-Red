package main

import (
	"strconv"
)

type Character struct {
	Name      string
	Race      string
	Level     int
	HealthMax int
	Health    int
	Inventory map[string]int
	Skills    []string
	Money int
	Equipment map[string]Equipment
	Exp int
	ExpCap int
	Dmg int
}

func (c *Character) init(name, race string, level, max_health, health, money int, inventory map[string]int, skills []string) {
	c.Name = name
	c.Race = race
	c.Level = level
	c.HealthMax = max_health
	c.Health = health
	c.Inventory = inventory
	c.Skills = skills
	c.Money = money
	c.Equipment = map[string]Equipment{
		"Head": Empty,
		"Body": Empty,
		"Foot": Empty,
		"Weapon": Empty,
	}
	c.ExpCap = 100
	c.Dmg = 10
}

func (c *Character) gainExp(qt int) {
	c.Exp += qt
	SlowPrint(" EXP +", strconv.Itoa(qt), " ")
	if c.Exp >= c.ExpCap {
		c.Level++
		c.Exp = c.ExpCap - c.Exp
		c.ExpCap += 50
		c.HealthMax += 10
		c.Health = c.HealthMax
		c.Dmg++
		SlowPrint(c.Name, " gain a level !!!\n")
		SlowPrint("+10 HP Max\n+1 base DMG\n")
		WaitEnter()
	}
	SlowPrint(" EXP: ", strconv.Itoa(c.Exp), "/", strconv.Itoa(c.ExpCap), "\n")
}

func (c Character) displaySheet() {
	SlowPrint("Character's sheet:\n")
	SlowPrint(" Name: ", c.Name, "\n")
	SlowPrint(" Race: ", c.Race, "\n")
	SlowPrint(" Health: ")
	c.printHealth()
	SlowPrint(" Money: ", strconv.Itoa(c.Money), "\n")
	SlowPrint(" Level: ", strconv.Itoa(c.Level), "\n")
	SlowPrint(" EXP: ", strconv.Itoa(c.Exp), "/", strconv.Itoa(c.ExpCap), "\n")
	SlowPrint(" Skills:\n")
	for _, skill := range c.Skills {
		SlowPrint("   -", skill, "\n")
	}
	SlowPrint(" Equipment:\n")
	for slot, stuff := range c.Equipment {
		SlowPrint("   -", slot, ": ", stuff.Name, "\n")
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
}

func (c Character) printHealth() string {
	healthString := strconv.Itoa(c.Health) + "/" + strconv.Itoa(c.HealthMax) + "\n"
	if c.Health <= (c.HealthMax / 3) {
		SlowPrint(Colorize(Red, healthString))
	} else if c.Health > (c.HealthMax/3)*2 {
		SlowPrint(Colorize(Green, healthString))
	} else {
		SlowPrint(Colorize(Yellow, healthString))
	}
	return healthString
}

func (c *Character) dead() bool {
	if c.Health <= 0 {
		SlowPrint(Colorize(Red, c.Name, " is dead ...\n"))
		c.Health = c.HealthMax / 2
		SlowPrint("Fortunately, a priest who passed by was able to resurrect you.\nDon't waste your second chance.\n")
		WaitEnter()
		MainMenu()
		return true
	}
	return false
}

func (c *Character) addToHealth(qt int) bool {
	c.Health += qt
	if c.Health > c.HealthMax {
		c.Health = c.HealthMax
	}
	return c.dead()
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
	var inv = map[string]int{
		"Life Potion": 4,
		"Adventurer Hat": 2,
		"Adventurer Cloak": 1,
		"Crown": 1,
	}
	Player.init(Colorize(Cyan, "Benjamin"), "Human", 1, 100, 100, 100, inv, []string{"Punch"})
}
