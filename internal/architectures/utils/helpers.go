package utils

import (
	"os"
	"path/filepath"

	"github.com/K31NER/gobootx/internal/docker"
)

// CreateDirs crea una lista de directorios dentro de una ruta base
func CreateDirs(basePath string, paths []string) error {
	for _, path := range paths {
		fullPath := filepath.Join(basePath, path)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

// CreateDockerFile genera un Dockerfile basado en el lenguaje seleccionado
func CreateDockerFile(basePath string, language string) error {
	dockerContent := ""

	switch language {
	case "Go":
		dockerContent = docker.GoDockerfile
	case "Python":
		dockerContent = docker.FastapiDockerfile
	}

	if dockerContent == "" {
		return nil
	}

	dockerPath := filepath.Join(basePath, "Dockerfile")
	return docker.CreateDockerfile(dockerPath, dockerContent)
}
