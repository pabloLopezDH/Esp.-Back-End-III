package main

import (
	"fmt"
)

/*
Ejercicio 1 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C: su salario es de $1.000 por hora.
Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría,
y que devuelva su salario.
*/

func main() {

	var (
		minTrabajados float64 = 1000
		categoria     string  = "A"
	)

	salario := calcularSalario(minTrabajados, categoria)
	fmt.Println("Calculo del sueldo: $", salario)

	salario2 := calcularSalario(minTrabajados, "B")
	fmt.Println("Calculo del sueldo: $", salario2)

	salario3 := calcularSalario(minTrabajados, "C")
	fmt.Println("Calculo del sueldo: $", salario3)
}

func calcularSalario(minMensuales float64, categoria string) float64 {

	var totalSalario float64
	var salario float64
	var porcentaje float64

	if categoria == "C" {
		totalSalario = minMensuales * 16.66
	} else if categoria == "B" {
		salario = minMensuales * 25
		porcentaje = (salario / 100) * 20
		totalSalario = porcentaje + salario
	} else if categoria == "A" {
		salario = minMensuales * 50
		porcentaje = (salario / 100) * 50
		totalSalario = porcentaje + salario
	}

	return totalSalario

}
