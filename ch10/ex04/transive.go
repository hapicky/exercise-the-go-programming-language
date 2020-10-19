package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ListResult struct {
	ImportPath string
	Deps       []string
}

func (result ListResult) depend(p string) bool {
	for _, dep := range result.Deps {
		if dep == p {
			return true
		}
	}
	return false
}

func main() {
	packages := os.Args[1:]
	if len(packages) == 0 {
		fmt.Fprintln(os.Stderr, "usage: transive package1 package2 ...")
		os.Exit(1)
	}

	results := make([]ListResult, len(packages))
	for i, p := range packages {
		ljson, _ := exec.Command("go", "list", "-json", p).Output()
		json.Unmarshal(ljson, &results[i])
	}

	// 依存している順に並べる
	var depends []ListResult
	for _, result := range results {
		// 最初のパッケージは無条件で追加
		if len(depends) == 0 {
			depends = append(depends, result)
			continue
		}
		// 先頭のパッケージに依存していたら先頭に挿入
		if result.depend(depends[0].ImportPath) {
			depends = append([]ListResult{result}, depends...)
			continue
		}
		// 末尾のパッケージから依存されていたら末尾に追加
		if depends[len(depends)-1].depend(result.ImportPath) {
			depends = append(depends, result)
			continue
		}
	}

	// 全てが依存関係にない場合は終了
	if len(packages) != len(depends) {
		os.Exit(0)
	}

	// go list -json ... の結果はパースできなかったので欲しい情報をフォーマットで出力して1行ずつパース
	format := "{\"ImportPath\": \"{{.ImportPath}}\", \"Deps\": [\"{{join .Deps \"\\\", \\\"\"}}\"]}"
	alljson, _ := exec.Command("go", "list", "-f", format, "...").Output()

	for _, j := range strings.Split(string(alljson), "\n") {
		var result ListResult
		json.Unmarshal([]byte(j), &result)
		// 先頭のパッケージに依存するものを報告
		if result.depend(depends[0].ImportPath) {
			fmt.Println(result.ImportPath)
		}
	}
}
