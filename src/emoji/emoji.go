package emoji

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const version = "12.0"

type Emoji struct {
	Version string `json:"version"`
}

func NewEmoji() Emoji {
	return Emoji{
		Version: version,
	}
}

func (e Emoji) ShortCodeList() []string {
	var keys []string
	for k := range MapEmoji {
		keys = append(keys, k)
	}
	return keys
}

func (e Emoji) CodePoints() []string {
	var values []string
	for _, v := range MapEmoji {
		vv := `\` + strings.Replace(v, "+", "000", -1)
		fmt.Println(vv)
		fmt.Println(strconv.QuoteToASCII(vv))
		values = append(values, vv)
	}
	fmt.Println("\U0001F47B")
	return values
}

func (e Emoji) Code(shortCode string) (interface{}, error) {
	return "", nil
}

func (e Emoji) UnicodeByShortCode(shortCode string) (interface{}, error) {
	return "", nil
}

func (e Emoji) UnicodeByCharacter(r rune) (interface{}, error) {
	return fmt.Printf("%#U\n", unicode.SimpleFold(r))
}

func (e Emoji) Exists(shortCode string) (bool, error) {
	return true, nil
}

func (e Emoji) Markdown(shortCode string) (string, error) {
	return "", nil
}

func (e Emoji) Search(shortCode string) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (e Emoji) Versions() (string, error) {
	if e.Version == "" {
		return "No version", nil
	}
	return e.Version, nil
}

func (e Emoji) Random() (map[string]string, error) {
	return nil, nil
}

func (e Emoji) RandomNumber(number int) (map[string]string, error) {
	if number == 0 {
		return e.Random()
	}
	return nil, nil
}
