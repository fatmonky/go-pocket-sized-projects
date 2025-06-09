package main

import "fmt"

type language string

func main() {
	greeting := greet("cn")
	fmt.Println(greeting)
}

var phrasebook = map[language]string{
	"en": "Hello, world!",
	"fr": "Bonjour, monde!",
	"cn": "你好！",
	"de": "Hallo, welt!",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
