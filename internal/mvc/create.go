package mvc

import (
	"os"
	"path/filepath"

	"github.com/K31NER/gobootx/internal/boot"
)

type MCVBooot struct{}

func (m MCVBooot) Name() string {
	return "MVC"
}

func (m MCVBooot) Run(basePath string) error {
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

	return nil
}


func init(){
	boot.Register("2",MCVBooot{})
}