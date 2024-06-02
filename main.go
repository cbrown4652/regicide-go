package main

import (
	"fmt"
	"os"

	"github.com/cbrown4652/regicide-go/game"
	"github.com/cbrown4652/regicide-go/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	gameState game.GameState
}

func initialModel() model {
	gs := game.GameState{
		Deck:       game.Deck{},
		PlayerHand: game.Hand{},
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
			m.gameState.PlayerHand.ToggleCardSelection(0)
		case "2":
			m.gameState.PlayerHand.ToggleCardSelection(1)
		case "3":
			m.gameState.PlayerHand.ToggleCardSelection(2)
		case "4":
			m.gameState.PlayerHand.ToggleCardSelection(3)
		case "5":
			m.gameState.PlayerHand.ToggleCardSelection(4)
		case "6":
			m.gameState.PlayerHand.ToggleCardSelection(5)
		case "7":
			m.gameState.PlayerHand.ToggleCardSelection(6)
		case "8":
			m.gameState.PlayerHand.ToggleCardSelection(7)
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
