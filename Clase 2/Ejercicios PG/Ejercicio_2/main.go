package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función
en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

func main() {

	//promedio, err := promedioAlumnos(2, 4, -5, 8, 2)
	promedio, err := promedioAlumnos(2, 4, 5, 8, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio del alumno es:", promedio)
	}
}

func promedioAlumnos(notas ...int) (int, error) {

	var resultado int
	var promedio int

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Hay numeros negativos en las calificaciones.")
		} else {
			resultado += value
		}
	}

	promedio = resultado / len(notas)

	return promedio, nil

}
