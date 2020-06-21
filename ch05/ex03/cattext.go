package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cattext: #{err}\n")
		os.Exit(1)
	}
	catText(doc)
}

func catText(n *html.Node) {
	if n.Type == html.ElementNode &&
		(n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		if text := strings.TrimSpace(n.Data); text != "" {
			fmt.Println(text)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		catText(c)
	}
}
