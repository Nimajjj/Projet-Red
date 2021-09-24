package main

import (
  "log"
  "os"
)


func NewCharacter() {
  os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
  var name, race string
  SlowPrint("Welcome into THE GAME!\n")
  Wait(1)
  name = takeName()
  SlowPrint("Ho sorry ! \nHello ", name, ", select a race for your character.\n")
  race = takeRace()
  SlowPrint("Now you are ", name, " the ", race, ".\n")

  maxHealth := 100
  switch race {
  case "Elf":
    maxHealth -= 20
  case "Dwarf":
    maxHealth += 20
  }

  Player.init(name, race, 1, maxHealth, (maxHealth / 2), map[string]int{"Life Potion":3}, []string{"Punch"})
}

func takeRace() string {
  race := ""
  SlowPrint("0 - Human | 1 - Elf | 2 - Dwarf\n")
  input := TakeIntInput()
  if !OutOfRange(input, 0, 2) {
    switch input {
    case 0:
      race = "Human"
    case 1:
      race = "Elf"
    case 2:
      race = "Dwarf"
    default:
      log.Fatal("Error: input passed !OutOfRange func")
    }
  return race
  } else {
    race = takeRace()
  }
  return race
}

func takeName() string {
  SlowPrint("...Erm, what is your name again?\n")
  name := TakeStrInput()
  //Capitalize name here
  return Capitalize(name)
}
