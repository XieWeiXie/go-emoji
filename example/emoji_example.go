package example

import (
	"fmt"
	"go-emoji/src/emoji"
)

func EmojiExample() {
	emo := emoji.NewEmoji()
	fmt.Println(emo.Version)
	fmt.Println(emo.Versions())

	// all short code , such as :mahjong:
	fmt.Println(emo.ShortCodeList())

	for _, i := range emo.ShortCodeList() {
		emo.Println(i)
	}
	for _, i := range emo.ShortCodeList() {
		emo.Println(i)
	}
	for _, i := range emo.ShortCodeList() {
		fmt.Println(emo.Exists(i))
	}
	for _, i := range emo.ShortCodeList() {
		emo.Sprintf("%#v", i)
	}
	for _, i := range emo.ShortCodeList() {
		emo.Sprint(i)
	}
	for _, i := range emo.ShortCodeList() {
		fmt.Println(emo.UnicodeByShortCode(i))
	}
}
