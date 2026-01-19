package clean

import (
	"os"
	"path/filepath"

	"github.com/K31NER/gobootx/internal/boot"
)

type CleanBoot struct{}

func (c CleanBoot) Name() string {
	return "Clean Architecture"
}

func (c CleanBoot) Run(basePath string) error {

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

	return nil
}

func init(){
	boot.Register("1", CleanBoot{})
}