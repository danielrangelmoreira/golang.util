package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the front page
	resp, err := http.Get("https://www.bet365.com")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	//<span class="gl-Participant_Name">Novo Hamburgo</span>
	// define a matcher
	/*matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.Span {
			return true //scrape.Attr(n, "class") == "gl-Participant_Name"
		}
		return false
	}*/
	// grab all articles and print them

	items := scrape.FindAllNested(root, scrape.ByTag(atom.Span))
	for _, item := range items {
		fmt.Printf("%q\n", item.Data)
		for _, a := range item.Attr {
			fmt.Printf("%q (%q)\t", a.Key, a.Val)
		}
		fmt.Println()
	}
}
