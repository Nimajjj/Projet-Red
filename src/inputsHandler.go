package main

import (
  "os"
  "strconv"
)


func TakeIntInput() int {
  SlowPrint(">>> ")
  str, _ := Reader.ReadString('\n')
  if str[:len(str)-2] == "exit" {
    os.Exit(0)
  }
  res, err := strconv.Atoi(str[:len(str)-2])
  if err != nil || res < 0 {
    return -1
  }
  return res
}


func TakeStrInput() string {
  SlowPrint(">>> ")
  str, _ := Reader.ReadString('\n')
  if str[:len(str)-2] == "exit" {
    os.Exit(0)
  }
  return str[:len(str)-2]
}

func WaitEnter() {
  SlowPrint("Press 'enter' to continue...")
  Reader.ReadString('\n')
}
