package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required, e.q. en, fr...")
	// lang := flag.String("lang", "en", "The required, e.q. en, fr...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"el": "Γεια σου κόσμε",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
