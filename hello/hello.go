package main

import "fmt"

const defaultPrefix = "Hello, "

const spanish = "Spanish"
const spanishPrefix = "Hola, "

const french = "French"
const frenchPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = defaultPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Jarjar", ""))
}
