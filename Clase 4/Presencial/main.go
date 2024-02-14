package main

import (
	"fmt"
	"math"
)

/*
Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en Go.
Para ello requieren una estructura Matrix que tenga los métodos:
Set: recibe una matriz de flotantes e inicializa los valores en la estructura Matrix.
Print: imprime por pantalla la matriz de una forma más visible (con los saltos de línea entre filas).
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del
ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func main() {
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 54, 65, 76, 87, 87},
		alto:    3,
		ancho:   3,
	}

	Matrix.Set(m)
	Matrix.Print(m)
	Matrix.Cuadratica(m)
}

func (m Matrix) Set() {
	if len(m.valores) != m.ancho*m.alto {
		fmt.Println("La cantidad de valores no coincide con las dimensiones especificadas.")
	}

}

func (m Matrix) Cuadratica() bool {
	if (m.alto == m.ancho) && m.alto != 0 {
		return true
	}
	return false
}

func (m Matrix) Max() float64 {
	max := -math.MaxFloat64
	for _, elemento := range m.valores {
		if elemento > max {
			max = elemento
		}
	}
	return max
}

func (m Matrix) Print() {
	if len(m.valores) == 0 {
		fmt.Println("La matriz está vacía")
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Println(m.valores[fila*m.ancho : fila*m.ancho+m.ancho])
	}
}
