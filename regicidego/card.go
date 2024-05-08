package regicidego

type Card struct {
	suit string
	rank string
}

func (c *Card) Suit() string {
	return c.suit
}

func (c *Card) Rank() string {
	return c.rank
}
