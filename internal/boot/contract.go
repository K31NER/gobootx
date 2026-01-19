package boot

type BootConfig struct {
	Language   string
	WithDocker bool
}

// Definimos el contrato que debem seguir los boots
type Bootschema interface {
	Name() string
	Run(basePath string, config BootConfig) error
}