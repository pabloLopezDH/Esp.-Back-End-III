package main

import (
	"fmt"
	"os"
)

/*
Ejercicio - Mayorista
Se necesita un software desarrollado en Go que imprima las estimaciones de lo que generarían,
los productos de cada categoría en un mayorista, en ventas de productos para el hogar.
Para eso se detallarán los pasos para conseguirlo:

	1- Funcionalidad para generar un archivo CSV, llamado categorías.csv.
	2- Cargar el archivo con las categorías. Ejemplo:
		001	Electrodomésticos	ListaProductos
		002	Muebles		ListaProductos
		003	Herramientas		ListaProductos
		004	Pinturas		ListaProductos
		005	Aberturas		ListaProductos
		006	Construcción		ListaProductos
		007	Automotor 		ListaProductos
		Etcétera…
	Elegir al menos tres categorías.

	3- Generar para cada una de estas categorías los productos. Estos tendrán como información:
Código
Nombre
PrecioActual
CantidadActual
	Insertar al menos cuatro productos.
	4- Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma
de todos los productos de cada una de las categorías.
	5- Imprimir todos los estimativos por consola. Ejemplo:
Categoría			EstimativoPorCategoría
Construcción				60.700
Pinturas 				40.500
Aberturas				55.300
TotalEstimativo 			156.500
*/

func main() {

	// Creo el archivo e inserto datos
	categorias := []byte("IdCategoria,Nombre,ListaProducto;001,Construcción,ListaProducto;002,Pinturas,ListaProducto;003,Aberturas,ListaProducto;")
	os.WriteFile("./categorías.csv", categorias, 0644)

	// Leo el archivo
	res, err := os.ReadFile("./categorías.csv")

	if err != nil {
		fmt.Println("Error en la lectura del archivo!")
	}

	datosString := string(res)
	fmt.Println(datosString)

}
