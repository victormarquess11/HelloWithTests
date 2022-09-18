package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const english = "English"
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	prefix := greetingPrefix(language)
	name = greetingPosfix(name, language)

	return prefix + name
}

func greetingPosfix(name string, language string) (posfix string) {

	if name == "" {
		switch language {
		default:
			name = "World"
		case english:
			name = "World"

		case french:
			name = "Monde"

		case spanish:
			name = "Mundo"
		}
	}

	return name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	default:
		prefix = greetingPrefix(english)
	case english:
		prefix = englishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
