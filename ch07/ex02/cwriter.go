package main

import (
	"fmt"
	"io"
	"strings"
)

type counter struct {
	writer io.Writer
	size   int64
}

func (c *counter) Write(p []byte) (int, error) {
	l := len(p)
	c.size += int64(l)
	return l, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := counter{writer: w}
	return &counter, &counter.size
}

func main() {
	var sb strings.Builder
	cw, size := CountingWriter(&sb)
	fmt.Fprint(cw, "I have a pen")
	fmt.Println(*size)
	fmt.Fprint(cw, "I have an apple")
	fmt.Println(*size)
}
