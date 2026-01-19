package tui

import (
	"os"

	"github.com/K31NER/gobootx/internal/boot"
	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	cursor int
}

func InitialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd{
	return  nil
}

// Captura los casos
func (m model) Update(msg tea.Msg)(tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "1":
			if len(boot.Schemas) > 0 {
				wd, _ := os.Getwd()
				boot.Schemas[0].Run(wd)
			}
			return  m, tea.Quit

		case "q" , "ctrl+c":
			return m, tea.Quit
		}
		
	}

	return m,nil
}

// Iniciamos la TUI
func (m model) View() string {
	return `
	GoBootX

	[1] Crear structura Clean Architecture
	[q] Salir
	`
}

