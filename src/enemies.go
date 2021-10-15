package main

type Enemy struct {
  Name string
  MaxHealth int
  Health int
  Dmg int
  CritChance int
  Level int
  LootTable []string
  Exp int
  id int
}

func (e *Enemy) initEnemy(name string, maxHealth, health, dmg, crit, level int, lootTable []string) {
  e.Name = name
  e.MaxHealth = maxHealth
  e.Health = health
  e.Dmg = dmg * ( 1 + ( level / 100 ) )
  e.CritChance = crit
  e.Level = level
  e.LootTable = lootTable
  e.Exp = 10 * ( (level + 1) / 2)
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
  goblinLootTable := []string{
    "Goblin Underwear",
    "Goblin tooth",
  }
  gnomLootTable := []string{
    "Four-leaf clover",
    "Gnom Skull",
    "Gnom spin",
  }
  Goblin.initEnemy(Colorize(Purple, "Goblin"), 40, 40, 6, 10, 2, goblinLootTable)
  Gnom.initEnemy(Colorize(Purple, "Gnom"), 25, 25, 3, 50, 1, gnomLootTable)
  P.initEnemy("Player", 1, 1, 1, 1, 1, []string{})
  P.id = -1
}
