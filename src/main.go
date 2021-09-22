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

  InitCharacter()

  fmt.Println()
  InitMenus()
  DisplayMenu(MainMenu)
}
