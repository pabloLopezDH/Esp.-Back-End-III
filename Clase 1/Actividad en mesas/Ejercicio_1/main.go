package main

import (
	"fmt"
)

/*
Ejercicio 1 - Listado de nombres
Una profesora de la universidad quiere tener un listado con todos sus estudiantes.
Es necesario crear una aplicación que contenga dicha lista. Sus estudiantes son:
Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
Luego de dos clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
sin modificar el código que escribiste inicialmente. El nombre de la nueva estudiante es Gabriela.
*/

func main() {

	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos",
		"Pedro", "Axel", "Alez", "Dolores", "Federico",
		"Hernan", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println(estudiantes)

	estudiantes = append(estudiantes, "Gabriela")

	fmt.Println(estudiantes)

}
