package emoji

import (
	"fmt"
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
