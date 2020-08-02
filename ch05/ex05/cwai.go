package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	if doc.Type == html.ElementNode && doc.Data == "img" {
		images += 1
	}
	if doc.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(doc.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words += 1
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		cwords, cimages := countWordsAndImages(c)
		words += cwords
		images += cimages
	}
	return
}

func main() {
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d words and %d images\n", words, images)
}
