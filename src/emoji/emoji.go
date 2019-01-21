package emoji

import (
	"fmt"
	"html"
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
		vv := strings.Replace(v, "U+", "", -1)
		values = append(values, vv)
	}
	return values
}

func (e Emoji) Map() map[string]int64 {
	var results = make(map[string]int64)
	for k, v := range MapEmoji {
		newV := strings.Replace(v, "U+", "", -1)
		//fmt.Println(newV, v)
		hex, _ := strconv.ParseInt(newV, 16, 64)
		results[k] = hex
	}
	return results
}

func (e Emoji) MapSingle(key string) int64 {
	allResults := e.Map()
	if v, ok := allResults[key]; ok {
		return v
	}
	return -1
}

func (e Emoji) compile(a *[]interface{}) {
	for i, x := range *a {
		if str, ok := x.(string); ok {
			(*a)[i] = e.MapSingle(str)
		}
	}
}

func (e Emoji) Print(a ...interface{}) (int, error) {
	var full string
	for _, i := range a {
		value := e.MapSingle(i.(string))
		var str string
		if value != 0 {
			str = html.UnescapeString("&#" + strconv.Itoa(int(value)) + ";")
			full += str
		} else {
			full += i.(string)
		}
	}
	return fmt.Print(full)
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
