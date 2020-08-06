package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"./links"
)

func crawl(base_url string) []string {
	fmt.Println(base_url)

	list, err := links.Extract(base_url)
	if err != nil {
		log.Print(err)
	}

	u, err := url.Parse(base_url)
	if err != nil {
		log.Fatal(err)
	}
	domain := u.Hostname()

	savepage(base_url, domain)
	for _, link := range list {
		fmt.Println(link)
		savepage(link, domain)
	}

	return list
}

func savepage(link, target_domain string) {
	u, err := url.Parse(link)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if u.Hostname() != target_domain {
		return
	}

	resp, err := http.Get(link)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	base_dir := "./pages/" + target_domain
	dir := filepath.Dir(base_dir + u.Path)
	os.Mkdir(dir, 0755)

	out_path := base_dir + u.Path
	if strings.HasSuffix(out_path, "/") {
		out_path += "index.html"
	}

	ioutil.WriteFile(out_path, body, 0644)
}

func main() {
	url := os.Args[1]
	crawl(url)
}
