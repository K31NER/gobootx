package cmd

import (
	"fmt"

	"github.com/K31NER/gobootx/tui"
	tea "github.com/charmbracelet/bubbletea"
)

// Ejecuta la TUI
func Execute() error {

	// inicializamos el modelo
	p := tea.NewProgram(tui.InitialModel())

	// Arrancamos la interfaz
	m, err := p.Run()
	if err != nil {
		return err
	}

	if tuiModel, ok := m.(tui.Model); ok && tuiModel.ExitMsg != "" {
		fmt.Println(tuiModel.ExitMsg)
	}

	return nil
}