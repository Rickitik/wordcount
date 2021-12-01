package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1]
	ss := strings.Split(argsWithoutProg, " ")
	if len(ss) == 1 && ss[0] == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(ss))
	}
}
