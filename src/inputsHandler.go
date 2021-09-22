package main

import (
  "fmt"
  "os"
  "strconv"
)

func TakeIntInput() int {
  fmt.Print(">>> ")
  str, _ := BR.ReadString('\n')
  if str[:len(str)-2] == "exit" {
    os.Exit(0)
  }
  res, err := strconv.Atoi(str[:len(str)-2])
  if err != nil || res < 0 {
    return -1
  }
  return res
}


func WaitEnter() {
  BR.ReadString('\n')
}
