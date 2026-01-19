package docker

import (
	"os"
	"path/filepath"
)

/*
Crea el Dockerfile

Parametros:
    - path: Ruta donde se generara el Dockerfile
	- content: Contenido que tendra el Dockerfile
*/
func CreateDockerfile(path string, content string) error {
	// Creamos el archivo
	if err:= os.MkdirAll(filepath.Dir(path), 0755);err != nil{
		return  err
	}

	// Escribimos en el Dockerfile
	return  os.WriteFile(path, []byte(content), 0644)
}
