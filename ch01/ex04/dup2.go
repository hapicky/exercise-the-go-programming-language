package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// 行の内容をキーとして、さらにファイル名をキーにするmapを持つmap
	// { "line content": { "foo.txt": true, "bar.txt": true } }
	lineFiles := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "STDIN", lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, lineFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			i := 0
			filenames := make([]string, len(lineFiles[line]))
			for k := range lineFiles[line] {
				filenames[i] = k
				i++
			}
			fmt.Printf("%d\t%s\t%s\n", n, filenames, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, lineFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		// ファイル名をキーとして保持させる
		_, ok := lineFiles[line]
		if !ok {
			lineFiles[line] = make(map[string]bool)
		}
		lineFiles[line][filename] = true
	}
}
