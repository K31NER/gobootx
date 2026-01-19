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

// Opciones afirmativas
var affirmativeOptions = map[string]bool{
	"y": true, "Y": true,
	"s": true, "S": true,
}

type Step int

const (
	StepSelectArch Step = iota
	StepSelectLang
	StepSelectDocker
)

type Model struct {
	textInput textinput.Model
	step      Step
	width     int
	height    int
	ExitMsg   string

	// Selections
	selectedArch string
	selectedLang string
	addDocker    bool
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Opción"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		textInput: ti,
		step:      StepSelectArch,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Captura los casos
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
    
	// Captuamos el tipo de mensaje
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
    
	// Capturamos las letras
	case tea.KeyMsg:
		// Capturamos el tipo de tecla
		switch msg.Type {
		
		// Cerrar programa
		case tea.KeyCtrlC, tea.KeyEsc:
			m.ExitMsg = "Saliendo..."
			return m, tea.Quit
		
		// Evaluar enter
		case tea.KeyEnter:
			val := m.textInput.Value() // Entrada del usuario

			if val == "q" {
				m.ExitMsg = "Saliendo..."
				return m, tea.Quit
			}

			switch m.step {
			
			// Arquitectura
			case StepSelectArch:
				if _, ok := boot.Schemas[val]; ok {
					m.selectedArch = val     // Valor seleccionado
					m.step = StepSelectLang  // Siguiente interfaz
					m.textInput.SetValue("") 
					m.textInput.Placeholder = "Seleccione un lenguaje"
				}
            
			// lenguaje
			case StepSelectLang:
				if lang, ok := boot.Languages[val]; ok {
					m.selectedLang = lang     // Valor seleccionado
					m.step = StepSelectDocker // Siguiente interfaz
					m.textInput.SetValue("")
					m.textInput.Placeholder = "y/n"
				}
            
			// Dockerfile
			case StepSelectDocker:

				// Validamos entrada
				if affirmativeOptions[val] {
					m.addDocker = true
				} else {
					m.addDocker = false
				}

				// Valida si el esquema esta agregado
				if schema, ok := boot.Schemas[m.selectedArch]; ok {
					wd, _ := os.Getwd() // Obtenemos e directorio

					// Pasamos los datos optenidos
					config := boot.BootConfig{
						Language:   m.selectedLang,
						WithDocker: m.addDocker,
					}

					// Ejecutamos el codigo pasando ruta y configuracion
					schema.Run(wd, config)
					m.ExitMsg = fmt.Sprintf("✅ Arquitectura %s (%s) creada con éxito!", schema.Name(), m.selectedLang)
					if m.addDocker {
						m.ExitMsg += " (con Dockerfile)"
					}
					return m, tea.Quit
				}
			}
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// Iniciamos la TUI
func (m Model) View() string {

	// Definimos el logo
	title := `
     ██████╗  ██████╗ ██████╗  ██████╗  ██████╗ ████████╗██╗  ██╗
    ██╔════╝ ██╔═══██╗██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝╚██╗██╔╝
    ██║  ███╗██║   ██║██████╔╝██║   ██║██║   ██║   ██║    ╚███╔╝ 
    ██║   ██║██║   ██║██╔══██╗██║   ██║██║   ██║   ██║    ██╔██╗ 
    ╚██████╔╝╚██████╔╝██████╔╝╚██████╔╝╚██████╔╝   ██║   ██╔╝ ██╗
     ╚═════╝  ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝    ╚═╝   ╚═╝  ╚═╝`
    
	// Estilos 
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

	var content string
    
	switch m.step {

	// Seleccionar la arquitectura
	case StepSelectArch:
		// Sacamos las arquitecturas disponibles
		var keys []string
		for k := range boot.Schemas {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		menu := "\n\tSeleccione una Arquitectura:\n\n"
		for _, k := range keys {
			menu += fmt.Sprintf("\t[%s] Crear estructura %s\n", k, boot.Schemas[k].Name())
		}
		menu += "\t[q] Salir\n"
		content = textStyle.Render(menu)
    
	// Seleccionar el lenguaje
	case StepSelectLang:
		var keys []string
		for k := range boot.Languages {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		menu := "\n\tSeleccione un Lenguaje:\n\n"
		for _, k := range keys {
			menu += fmt.Sprintf("\t[%s] %s\n", k, boot.Languages[k])
		}
		
		content = textStyle.Render(menu)
    
	// Seleccionar si tiene dockerfile
	case StepSelectDocker:
		menu := `
	¿Desea agregar Dockerfile?
	
	[y] Si
	[n] No
	`
		content = textStyle.Render(menu)
	}

	// Renderizamos la terminal con los estilos
	return style.Render(title) + "\n" + dirStyle.Render("Path: "+currentDir) + "\n" + content + "\n\n" + m.textInput.View()
}