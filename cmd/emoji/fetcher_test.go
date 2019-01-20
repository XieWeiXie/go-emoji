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
