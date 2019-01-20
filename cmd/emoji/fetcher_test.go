package emoji_fetch

import "testing"

func TestRun(tests *testing.T) {
	tt := []struct {
		Fetcher
	}{
		{
			*NewFetcher("https://emojipedia.org/emoji/"),
		},
	}
	for _, t := range tt {
		t.Fetcher.Run()
	}
}

func TestTableHandler(tests *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: "https://emojipedia.org/emoji/%F0%9F%98%80/",
		}, {
			url: "https://emojipedia.org/emoji/%F0%9F%91%8B%F0%9F%8F%BB/",
		},
	}
	for _, t := range tt {
		tableHandler(t.url)
	}
}
