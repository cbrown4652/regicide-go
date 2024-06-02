package game

import (
	"math/rand"
	"time"
)

type Deck struct {
	tavern  []Card
	castle  []Card
	discard []Card
}

func (d *Deck) NewDeck() {
	d.tavern = newTavern()
	d.castle = newCastle()
}

func (d *Deck) Tavern() []Card {
	return d.tavern
}

func (d *Deck) Castle() []Card {
	return d.castle
}

func (d *Deck) Discard() []Card {
	return d.discard
}

func newTavern() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	tavern := []Card{}

	for _, suit := range suits {
		for _, rank := range ranks {
			tavern = append(tavern, Card{Suit: suit, Rank: rank})
		}
	}

	tavern = shuffleCards(tavern)

	return tavern
}

func newCastle() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"K", "Q", "J"}
	castle := []Card{}

	for _, rank := range ranks {
		var rank_group []Card
		for _, suit := range suits {
			rank_group = append(rank_group, Card{Suit: suit, Rank: rank})
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

func (d *Deck) DealPlayerHand() Hand {
	const size = 8
	var hand Hand

	for range size {
		hand.Cards = append(hand.Cards, d.dealNextTavernCard())
	}

	return hand
}

func (d *Deck) dealNextTavernCard() *Card {
	var nextCard *Card

	if len(d.tavern) > 0 {
		nextCard = &d.tavern[len(d.tavern)-1]
		d.tavern = d.tavern[:len(d.tavern)-1]
	}

	return nextCard
}

func (d *Deck) DealNextEnemy() Enemy {
	nextCard := d.castle[len(d.castle)-1]
	d.castle = d.castle[:len(d.castle)-1]

	return NewEnemy(nextCard)
}
