package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}

func contains(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	var pre func(n *html.Node)

	pre = func(n *html.Node) {
		if n.Type == html.ElementNode && contains(n.Data, name) {
			elements = append(elements, n)
		}
	}

	forEachNode(doc, pre)
	return elements
}

func printNodes(nodes []*html.Node) {
	for _, node := range nodes {
		html.Render(os.Stdout, node)
		fmt.Println("")
	}
}

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ebtn: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "ebtn: %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "ebtn: %v\n", err)
		os.Exit(1)
	}

	images := ElementsByTagName(doc, "img")
	printNodes(images)

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	printNodes(headings)
}
