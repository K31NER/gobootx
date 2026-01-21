package clean

import (
	"github.com/K31NER/gobootx/internal/architectures/utils"
	"github.com/K31NER/gobootx/internal/boot"
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
    
	// Creamos las carpetas
    if err := utils.CreateDirs(basePath, paths); err != nil{
		return err
	}
    
	// Creamos docker file si aplica
	if config.WithDocker { 
		utils.CreateDockerFile(basePath, config.Language)
	}

	return nil
}

func init(){
	boot.Register("1", CleanBoot{})
}