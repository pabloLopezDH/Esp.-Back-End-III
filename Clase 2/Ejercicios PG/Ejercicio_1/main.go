package main

import (
	"fmt"
)

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en
cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará,
además, un 10 % (27 % en total).
*/

func main() {

	var sueldo float64 = 150000

	impuesto := impuestosSueldos(sueldo)

	fmt.Print("Impuesto a descontar: ", impuesto)
}

func impuestosSueldos(sueldo float64) float64 {

	var impuesto float64

	if sueldo > 50000 && sueldo <= 149000 {
		impuesto = (sueldo / 100) * 17
	} else if sueldo >= 150000 {
		impuesto = (sueldo / 100) * 27
	}

	return impuesto
}
