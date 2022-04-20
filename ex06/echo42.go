package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var (
		s = flag.String("s", " ", "separator")
		n = flag.Bool("n", false, "omit trailing newline")
	)
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if *n == false {
		fmt.Print("\n")
	}
}
