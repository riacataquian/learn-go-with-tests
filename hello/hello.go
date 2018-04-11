package main

import "fmt"

const defaultPrefix = "Hello, "

const spanish = "Spanish"
const spanishPrefix = "Hola, "

const french = "French"
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishPrefix + name
	}

	if language == spanish {
		return frenchPrefix + name
	}

	return defaultPrefix + name
}

func main() {
	fmt.Println(Hello("Jarjar", ""))
}
