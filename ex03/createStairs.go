package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max, _ := strconv.Atoi(os.Args[1])
	cnt := 0
	row := 1
	for {
		if cnt+row > max {
			break
		}
		for i := 0; i < row; i++ {
			fmt.Printf("*")
			cnt++
		}
		fmt.Print("\n")
		row++
	}
}
