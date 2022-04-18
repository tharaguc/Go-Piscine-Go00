package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
			fmt.Println(ERROR_MSG)
			os.Exit(1)
	}
		fmt.Print(s)
}

func calculationStr(args []string) (string, bool) {
	if len(args) != 3 {
		return "ERROR", false
	}
	x, x_err := strconv.Atoi(args[1])
	y, y_err := strconv.Atoi(args[2])
	if x_err != nil || y_err != nil {
		return "ERROR", false
	}
	sum:= "sum: " + strconv.Itoa(x + y) + "\n"
	difference := "difference: " + strconv.Itoa(x - y) + "\n"
	product := "product: " + strconv.Itoa(x * y) + "\n"
	quotient := "quotient: " + strconv.Itoa(x / y) + "\n"
	res := (sum + difference + product + quotient)
	return res, true
}