package main

import "fmt"

func main() {

	card := newCard()

	fmt.Println(card)
}

func newCard() string { //func newCard() string <-- This function will return a value of type 'string'
	return "Five of Diamonds"
}
