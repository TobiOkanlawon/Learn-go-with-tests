package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Allo"
)

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("Toby", "English"))
}
