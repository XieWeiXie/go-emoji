package emoji

import (
	"fmt"
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
