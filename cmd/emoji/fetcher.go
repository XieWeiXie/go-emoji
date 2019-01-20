package emoji_fetch

import (
	"fmt"
	"go-emoji/configs"
	"go-emoji/models"
	"go-emoji/pkg/database"
	"log"
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
	var urlChannel = make(chan string)
	var emojiChannel = make(chan models.Emoji)
	content, _ := DownLoader(f.RootUrl)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	doc.Find("table.emoji-list tr").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Text())
		td := selection.Find("td")
		href, _ := td.Eq(0).Find("a").Attr("href")
		fullHref := root + href
		go func(fullUrl string) {
			urlChannel <- fullUrl
		}(fullHref)
		go func(one models.Emoji) {
			emojiChannel <- one
		}(tableHandler(<-urlChannel))
		dataImport(<-emojiChannel, "postgres")
	})

}

func tableHandler(url string) models.Emoji {
	fmt.Println(url)
	content, _ := DownLoader(url)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	article := doc.Find("article")
	var one models.Emoji
	one.Url = url
	one.Title = article.Find("h1").Text()
	article.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {
		if i == 1 {
			one.CodePoints = selection.Find("td").Eq(1).Text()
		}
		if selection.Find("td").HasClass("shortcodes") {
			content := selection.Find("td").Eq(1).Text()
			fmt.Println(content)
			one.ShortCodes = content
		}
		if selection.Find("td").HasClass("tags") {
			one.Tags = strings.Split(selection.Find("td").Eq(1).Text(), ",")
		}

	})
	//fmt.Println(one)
	return one
}

func dataImport(one models.Emoji, dialect string) {
	log.Println("start insert data into db")
	if dialect == configs.POSTGRES {
		var versions models.Version
		if dbError := database.POSTGRESDIALECT.Where("version_name = ?", "12.0").First(&versions).Error; dbError != nil {
			log.Println("can not find version")
			return
		}
		var result models.Emoji
		result = one
		result.VersionID = versions.ID
		if dbError := database.POSTGRESDIALECT.Save(&result).Error; dbError != nil {
			log.Println("save record fail")
			return
		}
	}
}
