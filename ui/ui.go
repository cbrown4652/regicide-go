package ui

import (
	"strings"

	"github.com/cbrown4652/regicide-go/regicidego"
)

func RenderGameState(gs regicidego.GameState) string {
	var sb strings.Builder

	sb.WriteString("\nRegicide \n\n")

	sb.WriteString("Castle:\n")

	for _, card := range gs.Deck.Castle() {
		sb.WriteString("Suit: " + card.Suit() + ", Rank: " + card.Rank() + "\n")
	}

	sb.WriteString("Tavern:\n")

	// sb.WriteString("Length: " + strconv.Itoa(len(gs.Deck.Tavern())) + "\n")

	// for i, card := range gs.Deck.Tavern() {
	// 	if i > 28 {
	// 		sb.WriteString("Suit: " + card.Suit() + ", Rank: " + card.Rank() + "\n")
	// 	}
	// }

	sb.WriteString("Hand:\n")

	for _, card := range gs.PlayerHand.Cards {
		sb.WriteString("Suit: " + card.Suit() + ", Rank: " + card.Rank() + "\n")
	}

	sb.WriteString("Enemy:\n")

	sb.WriteString(gs.Enemy.Card.Rank() + " " + gs.Enemy.Card.Suit() + "\n")

	sb.WriteString("Tavern after dealing hand:\n")

	// sb.WriteString("Length: " + strconv.Itoa(len(gs.Deck.Tavern())) + "\n")

	// for i, card := range gs.Deck.Tavern() {
	// 	if i > 28 {
	// 		sb.WriteString("Suit: " + card.Suit() + ", Rank: " + card.Rank() + "\n")
	// 	}
	// }

	return sb.String()
}
