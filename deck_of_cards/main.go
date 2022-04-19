package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suit = [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}
var ranks = [13]string{"Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

const rows = 13
const columns = 4
const total = rows * columns

func shuffleDeck(workDeck []string) []string {
	shuffled := make([]string, len(workDeck))

	rand.Seed(time.Now().UTC().UnixNano())

	perm := rand.Perm(len(workDeck))

	for i, v := range perm {
		shuffled[v] = workDeck[i]
	}

	return shuffled
}

func main() {
	var initDeck []string

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			temp := ranks[i] + " of " + suit[j]
			initDeck = append(initDeck, temp)
		}
	}

	// sorted cards after being sorted
	for k, v := range initDeck {
		fmt.Printf("%d %s \n", k, v)
	}

	for i := 0; i < 50; i++ {
		fmt.Print("-")
	}

	shuffled_deck := shuffleDeck(initDeck)
	// after being shuffled
	for j, v := range shuffled_deck {
		fmt.Printf("%d %s \n", j, v)
	}
}
