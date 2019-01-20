package emoji_fetch

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	root = "https://emojipedia.org"
)

type Fetcher struct {
	RootUrl string `json:"root_url"`
}

func NewFetcher(url string) *Fetcher {
	return &Fetcher{
		RootUrl: url,
	}
}

func (f Fetcher) Run() {
	content, _ := DownLoader(f.RootUrl)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	doc.Find("table.emoji-list tr").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Text())
		td := selection.Find("td")
		href, _ := td.Eq(0).Find("a").Attr("href")
		fullHref := root + href
		fmt.Println(fullHref)

	})

}

func tableHandler(url string) {
	content, _ := DownLoader(url)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	doc.Find("article")
}
