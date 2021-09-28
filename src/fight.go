package main

import (
  "strconv"
  "time"
  "math/rand"
)

var Enemies []Enemy
var Turn []Enemy
var Logs string

func InitFight(enemies []Enemy, place string) {
  Clear()
  rand.Seed(time.Now().UnixNano())
  Enemies = enemies
  Logs = ""
  Turn = append(Turn, P)
  for i, enemy := range Enemies {
    enemy.id = i
    Enemies[i].id = i
    Turn = append(Turn, enemy)
  }
  SlowPrint(initIntroMsg(place), "\n")
  fightMenu(place)
}

func fightMenu(place string) {
  userMenu := Player.Name + ": "
  hpDisplayed := strconv.Itoa(Player.Health) + "/" + strconv.Itoa(Player.HealthMax) + "HP"
  if Player.Health <= Player.HealthMax/3 {
    hpDisplayed = Colorize(Red, hpDisplayed)
  } else if Player.Health >= (Player.HealthMax/3)*2 {
    hpDisplayed = Colorize(Green, hpDisplayed)
  } else {
    hpDisplayed = Colorize(Yellow, hpDisplayed)
  }
  userMenu += hpDisplayed + "\n"
  userMenu += "    0 - Attack\n    1 - Inventory\n"

  SlowPrint(initEnemyDisplayed())
  SlowPrint(Logs)
  SlowPrint(userMenu)

  turnHandler()
  fightMenu(place)
}

func turnHandler() {

  switch Turn[0] {
  case P:
    playerAction()
  default:
    enemyAction(Turn[0])
    WaitEnter()
  }
  Turn = ArrTurn(Turn)

  Clear()
}


func enemyAction(enemy Enemy) {
  chance := rand.Intn(100)
  if chance <= enemy.CritChance {
    Player.addToHealth((enemy.Dmg * 2) * -1)

  } else {
    Player.addToHealth(enemy.Dmg * -1)
    Logs += enemy.Name + " " + strconv.Itoa(enemy.id) + " attacked " + Player.Name + ": "
    Logs += Colorize(Red, strconv.Itoa(enemy.Dmg), "DMG\n")
  }
}


func playerAction() {
  var action, target int
  SlowPrint("Action ")
  action = TakeIntInput()
  if action == 0  && len(Enemies) > 1{
    SlowPrint("Target ")
    target = TakeIntInput()
  } else {
    target = 0
  }

  switch action {
  case 0:
    var dmg = -10
    Logs += Player.Name + " attacked " + Enemies[target].Name + " " + strconv.Itoa(Enemies[target].id) + ": "
    Logs += Colorize(Red, strconv.Itoa(dmg*-1), "DMG\n")
    if !Enemies[target].toHealth(dmg) {
      Turn = RemoveById(Turn, Enemies[target].id)
      Enemies = RemoveById(Enemies, Enemies[target].id)
      Logs += Player.Name + " killed " + Enemies[target].Name + " " + strconv.Itoa(Enemies[target].id) + ".\n"
    }
  case 1:
    Player.displayInventory()
  }
  if len(Enemies) == 0 {  // Fight ended
    Clear()
    SlowPrint("Congratulation ", Player.Name, "!\n")
    WaitEnter()
    MainMenu()
  }
}

func initEnemyDisplayed() string{  // please don't read this function, ty
  enemyDisplayed := ""
  for i, _ := range Enemies { // print '[id]': to fix
    var spaces = ""
    for j := 0; j < (len(Enemies[i].Name) - 8); j++ { // '8' == length of "| " + " Lvl.x"
      spaces += " "
    }
    enemyDisplayed += spaces + "["
    enemyDisplayed += strconv.Itoa(i)
    enemyDisplayed += "]" + spaces
  }
  enemyDisplayed += "\n"
  enemyDisplayed += " | "     // print ' | Name Lvl.lvl | ': ok
  for _, enemy := range Enemies {
    enemyDisplayed += enemy.Name
    enemyDisplayed += " Lvl." + strconv.Itoa(enemy.Level)
    enemyDisplayed += " | "
  }
  enemyDisplayed += "\n"
  for i, enemy := range Enemies {   // print 'HP/HPMAX': ok
    var spaces = ""
    for j := 0; j < (len(Enemies[i].Name) - 8)/2; j++{
      spaces += " "
    }
    enemyDisplayed += spaces + spaces
    hpDisplayed := strconv.Itoa(enemy.Health) + "/" + strconv.Itoa(enemy.MaxHealth) + "HP"
    if enemy.Health <= enemy.MaxHealth/3 {
      hpDisplayed = Colorize(Red, hpDisplayed)
    } else if enemy.Health >= (enemy.MaxHealth/3)*2 {
      hpDisplayed = Colorize(Green, hpDisplayed)
    } else {
      hpDisplayed = Colorize(Yellow, hpDisplayed)
    }
    enemyDisplayed += hpDisplayed
    enemyDisplayed += spaces
  }
  enemyDisplayed += "\n\n"
  return enemyDisplayed
}

func initIntroMsg(place string) string {
  introMsg := Player.Name + " began a fight with "
  if len(Enemies) >= 1 {
    for i, e := range Enemies {
      introMsg += "'" + e.Name + "'"
      if i != len(Enemies)-2 {
        introMsg += ", "
      } else {
        introMsg += " and "
      }
    }
  } else {
    introMsg += "'" + Enemies[0].Name + "'"
  }
  introMsg += "in '" + place + "'.\n"
  return introMsg
}

func RemoveById(arr []Enemy, id int) []Enemy {
  var res = []Enemy{}
  for _, val := range arr {
    if val.id != id {
      res = append(res, val)
    }
  }
  return res
}

func ArrTurn(arr []Enemy) []Enemy {
  res := make([]Enemy, len(arr))
  start := arr[0]
  for index := range arr {
    if index < len(arr) - 1{
      res[index] = arr[index + 1]
    } else {
      res[index] = start
    }
  }
  return res
}
