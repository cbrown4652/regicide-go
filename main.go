package main

import (
	"fmt"
	"os"

	"github.com/cbrown4652/regicide-go/regicidego"
	"github.com/cbrown4652/regicide-go/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	gameState regicidego.GameState
}

func initialModel() model {
	gs := regicidego.GameState{
		Deck:       regicidego.Deck{},
		PlayerHand: regicidego.Hand{},
	}
	gs.Deck.NewDeck()
	gs.PlayerHand = gs.Deck.DealPlayerHand()
	gs.Enemy = gs.Deck.DealNextEnemy()

	return model{gameState: gs}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.gameState.PlayerHand.Cards[0].Selected = true
		case "2":
			m.gameState.PlayerHand.Cards[1].Selected = true
		case "3":
			m.gameState.PlayerHand.Cards[2].Selected = true
		case "4":
			m.gameState.PlayerHand.Cards[3].Selected = true
		case "5":
			m.gameState.PlayerHand.Cards[4].Selected = true
		case "6":
			m.gameState.PlayerHand.Cards[5].Selected = true
		case "7":
			m.gameState.PlayerHand.Cards[6].Selected = true
		case "8":
			m.gameState.PlayerHand.Cards[7].Selected = true
		case "enter":
			m.gameState.PlayerPlaysCards(m.gameState.PlayerHand.Cards)
		}
	}

	return m, nil
}

func (m model) View() string {
	return ui.RenderGameState(m.gameState)
}

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
