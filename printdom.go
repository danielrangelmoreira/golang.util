package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func isAnchorElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func isTextNode(n *html.Node) bool {
	return n.Type == html.TextNode
}

func isHasOnlyOneChild(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild == n.LastChild
}

func getAttribute(n *html.Node, key string) (string, error) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}
	return "", errors.New(key + " not exist in attribute!")
}

func printRstLink(n *html.Node) {
	if !isHasOnlyOneChild(n) {
		fmt.Fprintf(os.Stderr, "Child number of anchor is not 1\n")
		return
	}

	if !isTextNode(n.FirstChild) {
		fmt.Fprintf(os.Stderr, "Child of anchor is not TextNode\n")
		return
	}

	text := strings.TrimSpace(n.FirstChild.Data)

	href, err := getAttribute(n, "href")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	rstLink := "`" + text + " <" + href + ">`__"
	fmt.Println(rstLink)
}

func traverse(n *html.Node) {
	if isAnchorElement(n) {
		printRstLink(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}

func parseCommandLineArguments() string {
	pPath := flag.String("input", "", "Path of HTML file to be processed")
	flag.Parse()
	path := *pPath
	if path == "" {
		fmt.Fprintf(os.Stderr, "Error: empty path!\n")
	}

	return path
}

func main() {
	inputFile := parseCommandLineArguments()

	fin, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Fail to Open File!\n")
		os.Exit(1)
		//panic("Fail to open " + inputFile)
	}
	defer fin.Close()

	doc, err := html.Parse(fin)
	if err != nil {
		panic("Fail to parse " + inputFile)
	}

	traverse(doc)
}
