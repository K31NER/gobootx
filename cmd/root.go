package cmd

import (
	"github.com/K31NER/gobootx/tui"
	tea "github.com/charmbracelet/bubbletea"
)

// Ejecuta la TUI
func Execute() error {

	// inicializamos el modelo
	p := tea.NewProgram(tui.InitialModel())

	// Arrancamos la interfaz
	_, err := p.Run()
	return err
}