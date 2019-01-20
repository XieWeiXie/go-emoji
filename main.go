package main

import (
	"fmt"
	"go-emoji/cmd"
	"io/ioutil"
	"net/http"
	"unicode"
)

func example() {
	fmt.Println("\U0001f4a4")
	fmt.Println("\U0001f34e")
	fmt.Println("\U0001f355")
	fmt.Println("\U0001F600")
	fmt.Println(rune('+'))
	fmt.Printf("%#U\n", unicode.SimpleFold('+'))
	fmt.Printf("%#U\n", unicode.SimpleFold('😀'))
	fmt.Println("\U0001F49E")
	fmt.Printf("%#U\n", unicode.SimpleFold('\U0001F49E'))
	fmt.Println(unicode.SimpleFold('\U0001F49E'))
	fmt.Printf("\U0001f485\U0001f3ff")
	res, err := http.Get("http://www.unicode.org/emoji/charts/full-emoji-list.html")
	fmt.Println(err)
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func main() {
	cmd.Execute()
}
