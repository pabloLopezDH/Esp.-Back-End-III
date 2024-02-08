package main

import "fmt"

/*
Ejercicio 1 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una
aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje.
La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/

func main() {

	var precio int = 4000
	var porcentaje int = 15

	descuento := (precio / 100) * porcentaje

	resultado := precio - descuento

	fmt.Println("Precio del producto:", precio, "$")
	fmt.Println("Descuento del", porcentaje, "%", "aplicado al producto:", descuento, "$")
	fmt.Println("Monto a pagar:", resultado, "$")

}
