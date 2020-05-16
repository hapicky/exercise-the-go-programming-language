package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}

	var sign string
	var bd string
	if s[:1] == "-" || s[:1] == "+" {
		sign = s[:1]
		bd = s[1:]
	} else {
		bd = s
	}

	return sign + commaNumber(bd)
}

func commaNumber(s string) string {
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
	values := []string{
		"",
		"12",
		"123",
		"1234",
		"12345",
		"123456",
		"1234567",
		"-12",
		"-123",
		"-1234",
		"-12345",
		"-123456",
		"+123456",
	}

	for i := 0; i < len(values); i++ {
		fmt.Printf("%q -> %q\n", values[i], comma(values[i]))
	}
}
