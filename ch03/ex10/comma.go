package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}

	m := l % 3
	if m == 0 {
		m = 3
	}

	var buf bytes.Buffer
	buf.WriteString(s[:m])

	for i := m; i < l; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i:i+3])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma(""))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}
