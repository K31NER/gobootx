package tui

import (
	"fmt"
	"os"
	"sort"

	"github.com/K31NER/gobootx/internal/boot"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type Model struct {
	textInput textinput.Model
	cursor    int
	width     int
	height    int
	ExitMsg   string
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Opción"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Captura los casos
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.ExitMsg = "Saliendo..."
			return m, tea.Quit
		case tea.KeyEnter:
			val := m.textInput.Value()

			if val == "q" {
				m.ExitMsg = "Saliendo..."
				return m, tea.Quit
			}
            
			// Obtenemos la opcion en base al numero de esta
			if schema, ok := boot.Schemas[val]; ok {
				wd, _ := os.Getwd()
				schema.Run(wd)
				m.ExitMsg = "✅ Arquitectura " + schema.Name() + " creada con éxito!"
				return m, tea.Quit
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// Iniciamos la TUI
func (m Model) View() string {
	title := `
     ██████╗  ██████╗ ██████╗  ██████╗  ██████╗ ████████╗██╗  ██╗
    ██╔════╝ ██╔═══██╗██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝╚██╗██╔╝
    ██║  ███╗██║   ██║██████╔╝██║   ██║██║   ██║   ██║    ╚███╔╝ 
    ██║   ██║██║   ██║██╔══██╗██║   ██║██║   ██║   ██║    ██╔██╗ 
    ╚██████╔╝╚██████╔╝██████╔╝╚██████╔╝╚██████╔╝   ██║   ██╔╝ ██╗
     ╚═════╝  ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝    ╚═╝   ╚═╝  ╚═╝`

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#87CEEB")).
		Align(lipgloss.Center)

	dirStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Align(lipgloss.Center)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Align(lipgloss.Left)

	if m.width > 0 {
		style = style.Width(m.width)
		dirStyle = dirStyle.Width(m.width)
	}

	currentDir, _ := os.Getwd()
    
	// Sacamos las arquitecturas disponibles
	var keys []string
	for k := range boot.Schemas {
		keys = append(keys, k)
	}

	// ordenamos
	sort.Strings(keys)
    
	// iniciamos el menu
	menu := "\n\tOpciones:\n\n"

	// Vamos agregando las opciones
	for _, k := range keys {
		menu += fmt.Sprintf("\t[%s] Crear estructura %s\n", k, boot.Schemas[k].Name())
	}

	// Definimos el de salida
	menu += "\t[q] Salir\n"
    
	// Renderizamos la terminal con los estilos
	return style.Render(title) + "\n" + dirStyle.Render("Path: "+currentDir) + "\n" + textStyle.Render(menu) + "\n\n" + m.textInput.View()
}

