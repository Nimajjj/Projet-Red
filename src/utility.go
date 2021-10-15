package main

import (
	"fmt"
	"time"
	"os"
	"log"
	"flag"
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

func SlowPrint(str ...string) { // bb 24/09/21 <3
	if Debug || Slow {
		for _, str_part := range str {
			fmt.Print(str_part)
		}
		return
	}

	for _, str_part := range str {
		for _, char := range str_part {
			fmt.Print(string(char))
			time.Sleep(5 * time.Nanosecond)
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

func DebugInit() {
	flag.BoolVar(&Debug, "debug", false, "Debug mode:\n Skip intro\n Disable slow print\n Init default character.")
	flag.StringVar(&BootState, "boot", "", "Boot on a specific menu:\n i: Inventory\n m: Merchant\n b: Blacksmith\n f: Training Fight")
	flag.BoolVar(&Slow, "s", false, "Disable slow print.")
	flag.BoolVar(&Intro, "i", false, "Skip intro\nInit default character")

	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
