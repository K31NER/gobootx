# GoBootX 🚀

> TUI interactiva para la creación automática de arquitecturas de desarrollo de software en Go.

![GoBootX Demo](img/terminal.png)

**GoBootX** es una herramienta de línea de comandos (CLI) con interfaz gráfica de terminal (TUI) diseñada para acelerar el inicio de nuevos proyectos. Permite generar automáticamente la estructura de carpetas y archivos base siguiendo los principios de **Clean Architecture**.

## ✨ Características

- **Interfaz Interactiva**: Navegación sencilla y estilizada gracias a [Bubble Tea](https://github.com/charmbracelet/bubbletea).
- **Clean Architecture**: Genera estructuras de proyecto robustas, escalables y estandarizadas listas para empezar a codificar.
- **Rápido y Eficiente**: Scaffolding instantáneo para tus nuevos microservicios o aplicaciones.

## 📦 Instalación

### Requisitos previos
- Go 1.18 o superior

### Opción 1: Go Install (Recomendado)

Si tienes Go instalado y configurado en tu PATH:

```bash
go install github.com/K31NER/gobootx@latest
```

### Opción 2: Compilar desde el código fuente

1. Clona el repositorio:
   ```bash
   git clone https://github.com/K31NER/gobootx.git
   cd gobootx
   ```

2. Instala las dependencias y compila:
   ```bash
   go mod tidy
   go build -o gobootx main.go
   ```

## 🚀 Uso

Navega al directorio donde quieres crear tu proyecto y ejecuta:

```bash
gobootx
```

Verás un menú interactivo:

1. Ingresa **"1"** y presiona Enter para generar la estructura de Clean Architecture.
2. Ingresa **"q"** o presiona `Esc` / `Ctrl+C` para salir.

### Estructura Generada (Ejemplo)

La herramienta creará una estructura similar a esta:

```text
.
├── cmd/
├── internal/
│   ├── domain/
│   ├── usecase/
│   └── infrastructure/
├── pkg/
└── ...
```

## 🛠️ Tecnologías

- **Lenguaje**: Go
- **TUI Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **Estilos**: [Lip Gloss](https://github.com/charmbracelet/lipgloss)
