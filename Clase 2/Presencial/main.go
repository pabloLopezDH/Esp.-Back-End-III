package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Calcular estadísticas
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
de calificaciones de los estudiantes de un curso. Requieren calcular los valores mínimo,
 máximo y promedio de sus calificaciones.
Para esto, se solicita generar una función que indique qué tipo de cálculo se
quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un
mensaje (en caso de que el cálculo no esté definido), que se le pueda pasar una
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
Veamos un ejemplo:

*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	minFun, err := operacion(minimum)
	if err != nil {
		log.Fatal(err)
	}

	maxFun, err := operacion(maximum)
	if err != nil {
		log.Fatal(err)
	}

	avgFun, err := operacion(average)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("El minimo es: ", minFun(3, 4, 7, 1, 9, 5, 6))
	fmt.Println("El maximo es: ", maxFun(3, 4, 7, 1, 9, 5, 6))
	fmt.Println("El promedio es: ", avgFun(3, 4, 7, 1, 9, 5, 6))

	// errorFunc, err := operacion("error")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("El error es: ", errorFunc(3, 4, 7, 1, 9, 5, 6))

	minFun2, maxFun2, avgFun2, err := operacionMulti()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nEl minimo es: %d | El maximo es: %d | El promedio es: %d",
		minFun2(3, 4, 7, 1, 9, 5, 6),
		maxFun2(3, 4, 7, 1, 9, 5, 6),
		avgFun2(3, 4, 7, 1, 9, 5, 6))

}

func operacionMulti() (func(...int) int, func(...int) int, func(...int) int, error) {
	return calculoMinimo, calculoMaximo, calculoPromedio, nil
}

func operacion(calculo string) (func(...int) int, error) {
	switch calculo {
	case minimum:
		return calculoMinimo, nil
	case maximum:
		return calculoMaximo, nil
	case average:
		return calculoPromedio, nil
	default:
		return nil, errors.New("El calculo no exite.")
	}
}

func calculoMinimo(values ...int) int {
	min := values[0]
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}

func calculoMaximo(values ...int) int {
	var max int
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

func calculoPromedio(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum / len(values)
}
