package main

import "fmt"

/*
Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a
clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.
*/

func main() {

	var (
		edad       int = 23
		antiguedad int = 2
		sueldo     int = 120000
	)

	var (
		edad_22años     int = 22
		antiguedad_1año int = 1
		sueldo_100mil   int = 100000
	)

	switch {
	case edad < edad_22años:
		fmt.Println("Debes tener 22 años para solicitar un credito.")
	case antiguedad < antiguedad_1año:
		fmt.Println("Tu antiguedad debe ser de un año para solicitar un credito.")
	case sueldo < sueldo_100mil:
		fmt.Println("Tu sueldo debe ser de 100 mil para solicitar un credito.")
	default:
		fmt.Println("Cumples con todos los requisitos para tu credito.")
	}

}
