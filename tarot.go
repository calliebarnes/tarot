package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	name        string
	description string
	meaning     string
}

func createDeck() []Card {
	deck := []Card{
		{"The Fool", "A happy-go-lucky pup, eager for adventure and new experiences.", "Embrace the enthusiasm and open-mindedness of the pup, and don't be afraid to take risks. Trust the journey and be open to the lessons life has in store for you."},
		{"The Magician", "A clever kitten juggling magical trinkets, creating a purr-fect show.", "Like the kitten, you have all the tools and resources you need to achieve your goals. Tap into your inner power and manifest your desires."},
		{"The High Priestess", "A wise, old owl perched between moonlit trees, sharing her knowledge through whispers.", "The owl encourages you to trust your intuition and seek the answers within. Pay attention to your dreams and subconscious thoughts."},
		{"The Empress", "A motherly cat surrounded by her kittens, basking in the sun.", "The Empress reminds you to nurture yourself and others. Take time to enjoy the simple pleasures in life."},
		{"The Emperor", "A stern dog sitting on a throne, surveying his kingdom.", "The Emperor reminds you to take charge of your life and set clear boundaries. Be assertive and confident in your decisions."},
	}

	return deck
}

func shuffleDeck(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func drawCards(deck []Card, n int) []Card {
	return deck[:n]
}

func main() {
	deck := createDeck()
	shuffleDeck(deck)
	numCardsToDraw := 1
	drawnCards := drawCards(deck, numCardsToDraw)

	for _, card := range drawnCards {
		fmt.Println(card.name)
		fmt.Println()
		fmt.Println(card.description)
		fmt.Println()
		fmt.Println(card.meaning)
	}
}
