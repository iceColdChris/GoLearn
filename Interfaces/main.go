package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {

}

//Omit the value cause its unused
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
