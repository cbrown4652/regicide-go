package regicidego

type Card struct {
	suit string
	rank string
}

func (c *Card) GetSuit() string {
	return c.suit
}

func (c *Card) GetRank() string {
	return c.rank
}
