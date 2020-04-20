package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(i) + ": " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
