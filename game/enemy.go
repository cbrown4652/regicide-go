package game

type Enemy struct {
	Card   Card
	Health int
	Attack int
}

func NewEnemy(c Card) Enemy {
	var health int
	var attack int

	switch c.Rank {
	case "J":
		health = 20
		attack = 10
	case "Q":
		health = 30
		attack = 15
	case "K":
		health = 40
		attack = 20
	default:
		health = 0
		attack = 0
	}

	return Enemy{Card: c, Health: health, Attack: attack}
}
