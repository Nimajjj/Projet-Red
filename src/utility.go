package main

import (
	"fmt"
	"time"
	"os"
)

const (
  Reset = "\033[0m"

  Red = "\033[31m"
  Green = "\033[32m"
  Yellow = "\033[33m"
  Blue = "\033[34m"
  Purple = "\033[35m"
  Cyan = "\033[36m"
  White = "\033[37m"
)

func SlowPrint(str ...string) { // Benjamin 24/09/21 <3
	for _, str_part := range str {
		for _, char := range str_part {
			fmt.Print(string(char))
			time.Sleep(40 * time.Nanosecond)
		}
	}
}

func Wait(t int) {
	time.Sleep(time.Duration(t) * time.Second)
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

func Colorize(color string, allStrings ...string) string {
  res := color
  for _, str := range allStrings {
    res += str
  }
  res += Reset
  return res
}

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
