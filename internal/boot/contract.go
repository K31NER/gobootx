package boot

// Definimos el contrato que debem seguir los boots
type Bootschema interface {
	Name() string
	Run(basePath string) error
}