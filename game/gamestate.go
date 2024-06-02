package game

import (
	"strconv"
	"strings"
)

type GameState struct {
	Deck       Deck
	PlayerHand Hand
	Enemy      Enemy
}

func (gs *GameState) PlayerPlaysCards(c []*Card) {
	var attack int
	for i := range c {
		if c[i].Selected {
			attack += rankToInt(c[i].Rank)
		}
	}

	gs.Enemy.Health -= attack
}

func rankToInt(rank string) int {
	rank = strings.ToUpper(rank)
	switch rank {
	case "A":
		return 1
	case "J", "Q", "K":
		return 10
	default:
		num, err := strconv.Atoi(rank)
		if err != nil {
			panic(err)
		}

		return num
	}
}
