package clean

import (
	"os"
	"path/filepath"

	"github.com/K31NER/gobootx/internal/boot"
	"github.com/K31NER/gobootx/internal/docker"
)

type CleanBoot struct{}

func (c CleanBoot) Name() string {
	return "Clean Architecture"
}

func (c CleanBoot) Run(basePath string, config boot.BootConfig) error {

	// Definimos las carpetas que vamos a crear
	paths := []string{
		"src/domain",
		"src/repository",
		"src/use_case",
		"src/infrastructure",
		"src/api",
		"src/config",
	}
    
	// Recorremos las carpetas
	for _, path := range paths {

		// Definimos la ruta de la carpeta
		fullPath := filepath.Join(basePath, path)

		// Validamos si se puede crear
		if err := os.MkdirAll(fullPath, 0755);err != nil{
			return err
		}
	}
    
	// Creamos docker file si aplica
	if config.WithDocker { 
		dockerContent := "" // Iniciamos la variable vacia

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
	boot.Register("1", CleanBoot{})
}