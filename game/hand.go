package game

type Hand struct {
	Cards []*Card
	limit int
}

func (h *Hand) ToggleCardSelection(i int) {
	h.Cards[i].Selected = !h.Cards[i].Selected
}
