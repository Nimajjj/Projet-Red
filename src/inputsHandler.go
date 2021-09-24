package main

import (
  "fmt"
  "os"
  "strconv"
)


func TakeIntInput() int {
  fmt.Print(">>> ")
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
  fmt.Print(">>> ")
  str, _ := Reader.ReadString('\n')
  if str[:len(str)-2] == "exit" {
    os.Exit(0)
  }
  return str[:len(str)-2]
}

func WaitEnter() {
  Reader.ReadString('\n')
}
