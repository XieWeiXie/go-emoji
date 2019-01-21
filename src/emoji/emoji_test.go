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

	emoji := NewEmoji()
	for _, i := range emoji.CodePoints() {
		hex, _ := strconv.ParseInt(i, 16, 64)
		str := html.UnescapeString("&#" + strconv.Itoa(int(hex)) + ";")
		fmt.Println(str)
	}

}

func TestMap(tests *testing.T) {
	emoji := NewEmoji()
	fmt.Println(emoji.Map())
}

func TestMapSinge(tests *testing.T) {
	emoji := NewEmoji()
	tt := []struct {
		value string
	}{
		{
			value: ":yum:",
		}, {
			value: ":anguished:",
		},
	}
	for _, t := range tt {
		fmt.Println(emoji.MapSingle(t.value))
	}
}

func TestPrint2(tests *testing.T) {
	tt := []struct {
		value string
	}{
		{
			value: ":yum:",
		}, {
			value: ":anguished:",
		},
	}
	emoji := NewEmoji()
	for _, t := range tt {
		emoji.Print(t.value)
		emoji.Println(t.value)
		emoji.Code(t.value)
	}
}

func TestEmojiCode(tests *testing.T) {
	tt := []struct {
		value string
	}{
		{
			value: ":wink:",
		},
	}
	emoji := NewEmoji()
	for _, t := range tt {
		fmt.Println(emoji.Code(t.value))
	}
}
