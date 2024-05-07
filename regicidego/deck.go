package regicidego

import (
	"math/rand"
	"time"
)

type Deck struct {
	tavern  []Card
	castle  []Card
	discard []Card
}

func (d *Deck) InitializeDeck() {
	d.tavern = initializeTavern()
	d.castle = initializeCastle()
}

func (d *Deck) GetTavern() []Card {
	return d.tavern
}

func (d *Deck) GetCastle() []Card {
	return d.castle
}

func (d *Deck) GetDiscard() []Card {
	return d.discard
}

func initializeTavern() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	tavern := []Card{}

	for _, suit := range suits {
		for _, rank := range ranks {
			tavern = append(tavern, Card{suit, rank})
		}
	}

	tavern = shuffleCards(tavern)

	return tavern
}

func initializeCastle() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"J", "Q", "K"}
	castle := []Card{}
	rank_group := []Card{}

	for _, rank := range ranks {
		rank_group = nil
		for _, suit := range suits {
			rank_group = append(rank_group, Card{suit, rank})
		}
		rank_group = shuffleCards(rank_group)
		castle = append(castle, rank_group...)
	}
	return castle
}

func shuffleCards(cards []Card) []Card {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}
