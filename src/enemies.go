package main

type Enemy struct {
  Name string
  MaxHealth int
  Health int
  Dmg int
  CritChance int
  Level int
  id int
}

func (e *Enemy) initEnemy(name string, maxHealth, health, dmg, crit, level int) {
  e.Name = name
  e.MaxHealth = maxHealth
  e.Health = health
  e.Dmg = dmg
  e.CritChance = crit
  e.Level = level
}

func (e *Enemy) toHealth(qt int) bool{
  e.Health += qt

	if e.Health > e.MaxHealth {
		e.Health = e.MaxHealth
	} else if e.Health <= 0 {
    return false  // dead
  }
  return true     // still alive
}

var P Enemy
var Goblin Enemy
var Gnom Enemy

func InitEnemies() {
  Goblin.initEnemy(Colorize(Purple, "Goblin"), 40, 40, 6, 10, 2)
  Gnom.initEnemy(Colorize(Purple, "Gnom"), 25, 25, 3, 50, 1)
  P.initEnemy("Player", 1, 1, 1, 1, 1)
    P.id = -1
}
