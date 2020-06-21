package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"sort"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countelements: #{err}\n")
		os.Exit(1)
	}

	count := make(map[string]int)
	countElements(count, doc)

	elms := make([]string, 0, len(count))
	for elm := range count {
		elms = append(elms, elm)
	}
	sort.Strings(elms)
	for _, elm := range elms {
		fmt.Printf("%s: %d\n", elm, count[elm])
	}
}

func countElements(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data] += 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElements(count, c)
	}
}
