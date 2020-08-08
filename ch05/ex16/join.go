package main

import (
	"fmt"
	"strings"
)

func join(sep string, elms ...string) string {
	return strings.Join(elms, sep)
}

func main() {
	fmt.Println(join(", ", "foo", "bar", "baz"))
}
