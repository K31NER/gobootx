package boot

// instanciamos el mapa de tipo bootschema
var Schemas = make(map[string]Bootschema)

var Languages = map[string]string{
	"1": "Python",
	"2": "Go",
}

// Funcion para registrar arquitecturas
func Register(key string, schema Bootschema) {
	Schemas[key] = schema
}