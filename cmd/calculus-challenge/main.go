package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"proj-aplicado-6/internal/game"
)

func main() {
	program := tea.NewProgram(game.New(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao iniciar o jogo: %v\n", err)
		os.Exit(1)
	}
}
