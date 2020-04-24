//A perfect shuffle, also known as a faro shuffle, splits a deck of cards into equal halves (there must be an
//even number of cards), then perfectly interleaves them. Eventually a series of perfect shuffles
//returns a deck to its original order. For instance, with a deck of 8 cards named (1 2 3 4 5 6 7 8),
//the first shuffle rearranges the cards to (1 5 2 6 3 7 4 8), the second shuffle rearranges the cards to
//(1 3 5 7 2 4 6 8), and the third shuffle restores the original order (1 2 3 4 5 6 7 8).
//
//Your task is to write a program that performs a perfect shuffle and use it to determine how many perfect
//shuffles are required to return an n-card deck to its original order; how many perfect shuffles are required
//for a standard 52-card deck? When you are finished, you are welcome to read or run a suggested solution,
//or to post your own solution or discuss the exercise in the comments below.
package main

import (
	"fmt"
)

func perfectShuffle(deck []int) []int {
	length := len(deck)

	//first we will split the deck into two half left and right
	leftDeck := make([]int, length/2)
	rightDeck := make([]int, length/2)
	copy(leftDeck, deck)
	copy(rightDeck, deck[length/2:])

	for i := 0; i < len(leftDeck); i++ {
		deck[i*2] = leftDeck[i]
		deck[(i*2)+1] = rightDeck[i]
	}
	return deck
}

// Compare the values of two slices
func compareSliceValues(sliceA []int, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for i := 0; i < len(sliceA); i++ {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}
	return true
}

func main() {
	//we will hard code the number of cards here
	numberOfCards := 50000

	//we will generate a deck of these cards
	//surly their is a better way to do this eg python deck = range(10)
	deck := make([]int, numberOfCards)
	for i := 0; i < numberOfCards; i++ {
		deck[i] = i + 1
	}

	origanl_deck := make([]int, numberOfCards)
	copy(origanl_deck, deck)

	fmt.Println("the deck we are starting with is:")
	fmt.Println(deck)
	fmt.Println("")

	toggle := false
	for i := 0; toggle == false; i++ {
		//fmt.Println(perfectShuffle(deck))
		perfectShuffle(deck)
		if compareSliceValues(deck, origanl_deck) {
			toggle = true
			fmt.Println(i + 1)
		}
	}

}
