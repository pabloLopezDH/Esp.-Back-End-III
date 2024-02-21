package main

import "fmt"

func main() {
	var a int
	a = 10

	println("Valor Inicial de a: ", a)
	fmt.Println("Espacio de memoria de a: ", &a)

	calc(&a)

	println("\nValor final de a: ", a)

}

func calc(a *int) {
	*a = 16
	fmt.Println("Nuevo valor de a: ", *a)
	fmt.Println("Espacio de memoria de a: ", a)
}
