package main

import "fmt"

const (
	italian = "Italian"
	french = "French"

	englishHelloPrefix = "Hello, "
	italianHelloPrefix = "Ciao, "
	frenchHelloPrefix = "Bonjour, "
)



func Hello(name, language string) string {
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case italian:
		prefix = italianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}