package main

import (
  "strconv"
  "time"
  "math/rand"
  "fmt"
)

var Enemies []Enemy
var Turn []Enemy
var Logs = []string{}
var T int = 1


func InitFight(enemies []Enemy, place string) {
  Clear()
  rand.Seed(time.Now().UnixNano())
  Enemies = enemies
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
  if Turn[0] == P {
    userMenu += "    0 - Attack\n    1 - Inventory\n"
  }

  fmt.Print(initEnemyDisplayed())

  fmt.Print("Turn: ", strconv.Itoa(T-1), "\n\n")

  if len(Logs) >= 10 {
    Logs = Logs[1:]
  }
  for i, log := range Logs {
    if i == len(Logs)-1 {
      SlowPrint(log, "\n")
    } else {
      fmt.Print(log, "\n")
    }
  }

  SlowPrint("\n", userMenu)

  turnHandler()
  fightMenu(place)
}


func turnHandler() {

  switch Turn[0] {
  case P:
    playerAction()
    T++
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
    l := enemy.Name + " " + strconv.Itoa(enemy.id) + " attacked " + Player.Name + ": "
    l += Colorize(Red, strconv.Itoa(enemy.Dmg*2), "DMG!!!")
    Logs = append(Logs, l)

  } else {
    Player.addToHealth(enemy.Dmg * -1)
    l := enemy.Name + " " + strconv.Itoa(enemy.id) + " attacked " + Player.Name + ": "
    l += Colorize(Red, strconv.Itoa(enemy.Dmg), "DMG")
    Logs = append(Logs, l)
  }
}


func playerAction() {
  var action, target int
  SlowPrint("Action ")
  action = TakeIntInput()
  if OutOfRange(action, 0, len(Enemies)-1) {
    playerAction()
  }
  if action == 0  && len(Enemies) > 1 {
    SlowPrint("Target ")
    target = TakeIntInput()
  } else {
    target = 0
  }

  switch action {
  case 0:
    var dmg = -10
    l := Player.Name + " attacked " + Enemies[target].Name + " " + strconv.Itoa(Enemies[target].id) + ": "
    l += Colorize(Red, strconv.Itoa(dmg*-1), "DMG")
    if !Enemies[target].toHealth(dmg) {
      l += "\n" + Player.Name + " killed " + Enemies[target].Name + " " + strconv.Itoa(Enemies[target].id) + "."
      Turn = RemoveById(Turn, Enemies[target].id)
      Enemies = RemoveById(Enemies, Enemies[target].id)
    }
    Logs = append(Logs, l)
  case 1:
    Player.displayInventory(true)
  }
  if len(Enemies) == 0 {  // Fight WIN !
    Clear()
    SlowPrint("Congratulation ", Player.Name, "!\n")

    Enemies = []Enemy{}
    Turn = []Enemy{}
    Logs = []string{}
    T = 1

    WaitEnter()
    MainMenu()
  }
}


func initEnemyDisplayed() string{  // please don't read this function, it's only for format displayed info
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
