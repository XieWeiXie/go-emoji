package emoji

import (
	"fmt"
	"testing"
)

var emoji = NewEmoji()

func TestShortCodeList(tests *testing.T) {
	fmt.Println(emoji.ShortCodeList())
}

func TestCodePoints(tests *testing.T) {
	fmt.Println(emoji.CodePoints())
}

func TestMap(tests *testing.T) {
	fmt.Println(emoji.Map())
}

func TestMapSingle(tests *testing.T) {
	for k, _ := range MapEmoji {
		fmt.Println(emoji.MapSingle(k))
	}
}

func TestCompile(tests *testing.T) {
	for k, _ := range MapEmoji {
		fmt.Println(emoji.Compile(k))
	}
}

func TestPrint(tests *testing.T) {
	for _, i := range emoji.ShortCodeList() {
		emoji.Print(i)
	}
}

func TestPrintln(t *testing.T) {
	for _, i := range emoji.ShortCodeList() {
		emoji.Println(i)
	}
}

func TestPrintf(t *testing.T) {
	for _, i := range emoji.ShortCodeList() {
		emoji.Printf("%s\n", i)
	}
}

func TestCode(tests *testing.T) {
	for _, i := range emoji.ShortCodeList() {
		fmt.Println(emoji.Code(i))
	}
}

func TestEmoji_UnicodeByShortCode(t *testing.T) {
	for _, i := range emoji.ShortCodeList() {
		fmt.Println(emoji.UnicodeByShortCode(i))
	}
}

func TestEmoji_Versions(t *testing.T) {
	fmt.Println(emoji.Version)
}

func TestEmoji_Exists(t *testing.T) {
	for _, i := range emoji.ShortCodeList() {
		fmt.Println(emoji.Exists(i))
	}
}
