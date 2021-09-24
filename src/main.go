package main

import (
  "bufio"
  "log"
  "os"
)

var Reader *bufio.Reader
var Player Character


func main() {
  var testMode = false

  log.SetFlags(log.LstdFlags | log.Lshortfile)
  Reader = bufio.NewReader(os.Stdin)

  if testMode {
    InitDefaultCharacter()
  } else {
    NewCharacter()
    WaitEnter()
  }

  MainMenu()
}


func OutOfRange(i, min, max int) bool {
  if i >= min && i <= max {
    return false
  }
  return true
}

func IsStrInArray(str string, arr []string) bool {
  for _, val := range arr {
    if val == str {
      return true
    }
  }
  return false
}
