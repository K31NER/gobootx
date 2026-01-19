package mvc

import (
	"os"
	"path/filepath"

	"github.com/K31NER/gobootx/internal/boot"
	"github.com/K31NER/gobootx/internal/docker"
)

type MCVBooot struct{}

func (m MCVBooot) Name() string {
	return "MVC"
}

func (m MCVBooot) Run(basePath string, config boot.BootConfig ) error {

	// Definimos las carpetas a crear
	paths := []string{
		"src/model",
		"src/view",
		"src/controller",
		"src/config",
	}

	for _, path := range paths {

		// Definimos la ruta
		fullPath := filepath.Join(basePath, path)

		// Creamos la carpeta
		if err := os.MkdirAll(fullPath,0755); err != nil{
			return  err
		}
	}
    
	// Creamos docker file si aplica
	if config.WithDocker {
		dockerContent := "" // Inicializamos la varibale

		// Evaluamos el lenguaje
		switch config.Language {
		case "Go":
			dockerContent = docker.GoDockerfile
		case "Python":
			dockerContent = docker.FastapiDockerfile
		}
        
		// Creamos el dockerfile
		if dockerContent != "" {
			dockerPath := filepath.Join(basePath, "src/Dockerfile")
			return docker.CreateDockerfile(dockerPath, dockerContent)
		}
	}

	return nil
}


func init(){
	boot.Register("2",MCVBooot{})
}