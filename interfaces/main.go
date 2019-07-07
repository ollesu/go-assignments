package main

import "fmt"

// behaviours common to all bots: they all have getGreeting func that returns a string
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// custom logic
func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
