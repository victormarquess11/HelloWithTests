package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const english = "English"
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	prefix := englishHelloPrefix

	switch language {
	case english:
		prefix = englishHelloPrefix
		if name == "" {
			name = "World"
		}

	case french:
		prefix = frenchHelloPrefix
		if name == "" {
			name = "Monde"
		}

	case spanish:
		prefix = spanishHelloPrefix
		if name == "" {
			name = "Mundo"
		}
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
