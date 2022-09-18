package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	prefix := englishHelloPrefix
	if language == "English" {
		prefix = englishHelloPrefix
		if name == "" {
			name = "World"
		}
	}

	if language == spanish {
		prefix = spanishHelloPrefix
		if name == "" {
			name = "Mundo"
		}
	}

	if language == french {
		prefix = frenchHelloPrefix
		if name == "" {
			name = "Monde"
		}
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
