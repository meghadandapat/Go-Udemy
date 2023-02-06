package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//******************CUSTOM TYPE DECLARATIONS*************
//GO is not a Object Oriented lang
//Create a new type of 'deck' whhicch is a slice of strings
//the new deck type extends the behaviour of a slice of string

type deck []string

//returns a value of type deck
//requires no receiver

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardValues {
		for _, value := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}

	}
	return cards
}

// any variable of type deck now gets access to the print method
// d is reference to the actual copy of deck type variable, here cards
// this is called receiver functions, think of it as method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// multiple return values
//when to use parameters vs when to set up a receiver?->discussed later

// parameter
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// reference
func (d deck) deal1(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// /**********TYPE CONVERSION**********
// dect to a slice of string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	//ReadFile returns two values hence stored in two vars
	bs, err := ioutil.ReadFile(filename)
	//if file is read correctly then value of err will be nil otherwise it will be populated
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // quit the program with given status code. non-zero value in exit means error has occured
	}
	//type conversion -> byte sice to a slice of string
	s := strings.Split(string(bs), ",")
	//type conversion -> slice of string to deck
	return deck(s)

}

func (d deck) shuffle() {
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	for i := range d {
		// newPosition := r.Intn(len(d))
		newPosition := rand.Intn(len(d))            //random number
		d[i], d[newPosition] = d[newPosition], d[i] //swap number syntax

	}
}
