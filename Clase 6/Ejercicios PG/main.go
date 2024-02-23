package main

import (
	"os"
)

/*
Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto con la información de productos comprados,
separados por punto y coma (CSV).
Este archivo debe tener el ID del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

func main() {
	head := []byte("id,precio,cantidad;\n10,800,10;\n12,1230,589;\n")
	os.WriteFile("./data.csv", head, 0644)
}
