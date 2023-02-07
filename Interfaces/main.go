package main

import "fmt"

//interface: reuse code logic for different types

type bot interface {
	//if you are a type with a function called "getGreeting" and you "return a type of string", then you are now also of the type "bot"
	//English bot, we look at Spanish bot and both of them have a function called getGreeting and they both return a string. So we can now imagine that English bot and Spanish bot are now also of type bot.

	//interface syntax should list out all the functions with their parameters and return types
	getGreeting() string
}

//To say that a type satisfies an interface means that the type implements all of the functions contained in interface definition

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(b bot) {
	//We put together the new function called print greeting that accepted something of type bot.
	//we passed in English bot in Spanish bot to them, but English bot and Spanish bot are also members of type bot.
	fmt.Println(b.getGreeting())
}

// Note:same function name with different receiver type can exist but same func name and same receiver with different parameters cannot exists

func (eb englishBot) getGreeting() string {
	//custom logic for generating english greeting
	return "Hi There"
}
func (sb spanishBot) getGreeting() string {
	//custom logic for generating spanish greeting
	return "Hola"
}

//***********Concrete Type v/s interface type***********
/*
Concrete type: map, struct, int, string, englishBot ->can directly create values of these type
Interface type: bot ->cannot dirextly createvalues out of these type

1. Interfaces are not replacement of generic types
2. Interfaces are implicit ie we do not have to exlicitly mentiion englishBot is of type BOT
3. Interfaces are a contract to help us manage types
*/
