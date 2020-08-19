package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("this is a pen."))
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("this\nis\nit"))
	fmt.Println(lc)
}
