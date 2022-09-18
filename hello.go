package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	prefix := englishHelloPrefix
	if language == "English" {
		prefix = englishHelloPrefix
		if name == "" {
			name = "World"
		}
	}

	if language == "Spanish" {
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
