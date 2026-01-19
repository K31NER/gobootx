package boot

// instanciamos el slice de tipo bootschema
var Schemas = []Bootschema{}

// Funcion para registrar arquitecturas
func Register(schema Bootschema) {
	Schemas = append(Schemas, schema)
}