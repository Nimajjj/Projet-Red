package main

import (
  "time"
  "fmt"
)

func SlowPrint(str ...string) { // Benjamin 24/09/21 <3
  for _, str_part := range str {
    for _, char := range str_part {
      fmt.Print(string(char))
      time.Sleep(60_000_000 * time.Nanosecond)
    }
  }
}

func Wait(t int) {
  time.Sleep(time.Duration(t) * time.Second)
}
