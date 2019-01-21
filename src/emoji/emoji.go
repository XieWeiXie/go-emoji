package emoji

import (
	"errors"
	"fmt"
	"html"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func htmlHandler(value int) string {
	return html.UnescapeString("&#" + strconv.Itoa(value) + ";")
}

func (e Emoji) Compile(a ...interface{}) string {
	var full string
	for _, i := range a {
		value := e.MapSingle(i.(string))
		var str string
		if value != 0 {
			str = htmlHandler(int(value))
			full += str
		} else {
			full += i.(string)
		}
	}
	return full
}

func (e Emoji) Print(a ...interface{}) (int, error) {
	value := e.Compile(a...)
	return fmt.Print(value)
}

func (e Emoji) Println(a ...interface{}) (int, error) {
	value := e.Compile(a...)
	return fmt.Println(value)
}

func (e Emoji) Printf(format string, a ...interface{}) (int, error) {
	value := e.Compile(a...)
	return fmt.Printf(format, value)
}

func (e Emoji) Fprint(w io.Writer, a ...interface{}) (int, error) {
	value := e.Compile(a...)
	return fmt.Fprint(w, value)
}

func (e Emoji) Fprintln(w io.Writer, a ...interface{}) (int, error) {
	value := e.Compile(a...)
	return fmt.Fprintln(w, value)
}

func (e Emoji) Fprintf(w io.Writer, a ...interface{}) (int, error) {
	value := e.Compile(a...)
	return fmt.Fprintf(w, value)
}

func (e Emoji) Sprint(a ...interface{}) string {
	value := e.Compile(a...)
	return fmt.Sprintf(value)
}

func (e Emoji) Sprintf(format string, a ...interface{}) string {
	value := e.Compile(a...)
	return fmt.Sprintf(format, value)
}

func (e Emoji) Errorf(format string, a ...interface{}) error {
	value := e.Compile(a...)
	return errors.New(e.Sprintf(format, value))
}

func (e Emoji) Code(shortCode string) map[string]string {
	var code = make(map[string]string)
	code[shortCode] = e.Compile(shortCode)
	return code
}

func (e Emoji) UnicodeByShortCode(shortCode string) map[string]string {
	var code = make(map[string]string)
	if v, ok := MapEmoji[shortCode]; ok {
		code[shortCode] = v
	}
	return code
}

func (e Emoji) Exists(shortCode string) bool {
	if _, ok := MapEmoji[shortCode]; !ok {
		return false
	}
	return true
}

func (e Emoji) Versions() string {
	if e.Version == "" {
		return "No version"
	}
	return e.Version
}

func (e Emoji) Length() int {
	if MapEmoji == nil {
		return 0
	}
	return len(MapEmoji)
}

func (e Emoji) Random() map[string]string {
	length := e.Length()
	var code = make(map[string]string)
	rand.Seed(time.Now().Unix())
	key := e.ShortCodeList()[rand.Intn(length)]
	code[key] = e.Compile(MapEmoji[key])
	return code
}
