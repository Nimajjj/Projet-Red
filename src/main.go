package main

import (
  "fmt"
  "bufio"
  "os"
)

var BR *bufio.Reader
var Player Character

func main() {
  BR = bufio.NewReader(os.Stdin)
  InitVar()

  InitCharacter()
  InitItems()

  fmt.Println()
  InitMenus()
  DisplayMenu(MainMenu)
}
