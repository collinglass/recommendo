package main

import (
	"github.com/crosbymichael/feedparser"
	"net/http"
	"regexp"
)

func getwordcounts(url string) (string, map[string]int) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	feed, err := feedparser.ParseAtomFeed(r.Body)
	if err != nil {
		panic(err)
	}
	wc := make(map[string]int)

	fmt.Printf("Feed: %s\nTitle: %s\n", feed.Link[0].Href, feed.Title)

	for _, e := range feed.Entry {
		words := getwords(e.Title + ' ' + e.Summary)
		for word in words {
			wc[word] += 1
		}
	}
	return feed.Title, wc
}

func getwords(html string) {
	txt := regexp.compile(r'<[^>]+>').sub('',html)
	fmt.println(txt)
}

func main() {

}
