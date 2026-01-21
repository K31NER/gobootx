package mvc

import (
	"github.com/K31NER/gobootx/internal/architectures/utils"
	"github.com/K31NER/gobootx/internal/boot"
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

	// Creamos las carpetas usando utils
	if err := utils.CreateDirs(basePath, paths); err != nil {
		return err
	}
    
	// Creamos docker file si aplica
	if config.WithDocker {
		utils.CreateDockerFile(basePath, config.Language)
	}

	return nil
}


func init(){
	boot.Register("2",MCVBooot{})
}