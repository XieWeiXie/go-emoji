package emoji

import (
	"fmt"
	"html"
	"strconv"
	"testing"
)

func TestEmojiShortCodeList(tests *testing.T) {
	emoji := NewEmoji()
	fmt.Println(emoji.ShortCodeList())
}

func TestEmojiCodePoints(tests *testing.T) {
	emoji := NewEmoji()
	fmt.Println(emoji.CodePoints())
}

func TestPrint(tests *testing.T) {
	tt := []struct {
		value string
	}{
		{
			value: "U+1F616 ",
		},
	}
	for _, t := range tt {
		//input := bytes.NewBufferString(t.value)
		fmt.Println(strconv.Unquote(t.value))
	}
	emoji := NewEmoji()
	for _, i := range emoji.CodePoints() {
		r := []rune(i)
		for _, ii := range r {
			html.UnescapeString("&#" + strconv.Itoa(int(ii)) + ";")
			//fmt.Println(str)

		}

	}
}
