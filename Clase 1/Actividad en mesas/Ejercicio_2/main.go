package main

import (
	"fmt"
)

/*
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente map, debemos imprimir la edad de Benjamin.

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado, también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del map.
*/

func main() {

	var empleados = map[string]int{
		"Benjamin": 20,
		"Nahuel":   30,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}

	fmt.Println("\nEdad de Benjamin:", empleados["Benjamin"])

	for key, value := range empleados {
		if value >= 21 {
			fmt.Println("Personas que tienen 21 años o mas en la empresa:", key, "con", value, "años.")
		}
	}

	fmt.Println("\nAgregando el Empleado con nombre Federico")
	empleados["Federico"] = 25

	fmt.Println("\nEliminando el Empleado con nombre Pedro")
	delete(empleados, "Pedro")

	fmt.Println(empleados)

}
