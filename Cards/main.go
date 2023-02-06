package main

//**********VARIABLE DECLARATIONS*******************
/*
func main() {
	// statically typed lang
	// var card string = "Ace of Spades"
	card := "Ace of Spades" //defining new variable
	// for reassigning new val only =
	fmt.Println(card)
}

//variable can be defined outside main but not initialized

//basic go types
//bool, string, int, float64
*/

//**********FUNCTIONS AND RETURN TYPES*******************
/*
func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of diamonds"
}

//Files in the same package can freely call functions defined in other files.
*/

// *******************ARRAYS AND SLICES******************************

/*
// arrays -> fixed size ists
// slices -> can grow and shrink, list with spl properties
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") //append does not modify the old var instead it returns a new slice that we assign back to var
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of diamonds"
}
*/

func main() {
	//use custom type deck
	// cards := deck{"Ace of Diamonds", newCard()}

	cards := newDeck()
	// cards := newDeckFromFile("My_cards")
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("My_cards")

}
