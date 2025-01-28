package main

import "fmt"

const italian = "Italian"
const french = "French"
const englishHelloPrefix = "Hello, "
const italianHelloPrefix = "Ciao, "
const frenchHelloPrefix = "Bonjour, "


func Hello(name, language string) string {
	if name == ""{
		name = "World"
	}

	if language == italian{
		return italianHelloPrefix + name
	}

	if language == french{
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}