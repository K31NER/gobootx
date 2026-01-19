package boot

// instanciamos el mapa de tipo bootschema
var Schemas = make(map[string]Bootschema)

// Funcion para registrar arquitecturas
func Register(key string, schema Bootschema) {
	Schemas[key] = schema
}