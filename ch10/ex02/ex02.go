package main

import (
	"fmt"
	"os"

	"./areader"
	_ "./areader/tar"
	_ "./areader/zip"
)

// コマンドライン引数にアーカイブのパスを受け、そのアーカイブに含まれるファイル一覧を出力します
func main() {
	file := os.Args[1]

	names, err := areader.Filenames(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}
