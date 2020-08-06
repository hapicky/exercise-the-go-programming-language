package main

import (
	"fmt"
	"log"
	"os"

	"./links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	url := os.Args[1]
	links := crawl(url)

	for _, link := range links {
		fmt.Println(link)
	}
}
