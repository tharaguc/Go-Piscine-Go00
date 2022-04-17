package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument")
		return
	}
	adress := os.Args[1]
	if len(adress) <= 256 {
		regex := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)
		match := regex.MatchString(adress)
		if match {
			fmt.Printf("%v is a valid email adress.\n", adress)
			return
		}
	}
	fmt.Printf("%v is NOT a valid email address.\n", adress)
	return
}