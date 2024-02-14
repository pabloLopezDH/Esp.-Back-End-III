package main

import (
	"fmt"
)

/*
Ejercicio
Crear un programa que cumpla los siguientes puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el
slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá
imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto
correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definidos desde main().
*/

func main() {
	p := Product{3, "pantalon", "Pantalon militar", "Indumentaria", 4.0}
	fmt.Println("Productos antes de guardar uno nuevo:")
	p.GetAll()
	p.Save()
	fmt.Println("Productos después de guardar uno nuevo:")
	p.GetAll()

	idBuscado := 2
	if productFound, err := getByID(idBuscado); err != nil {
		fmt.Printf("Buscando el producto con el ID=%d. Error=%v", idBuscado, err)
	} else {
		fmt.Printf("Buscando el producto con el ID = %d. Encontrado= %v", idBuscado, productFound)
	}

}

var (
	Products = []Product{
		{1, "zapatilla", "Zapatillas deportivas", "Calzado", 40.0},
		{2, "remera", "Remeras básicas", "Indumentaria", 5.0},
	}
)

type Product struct {
	ID                          int
	Name, Description, Category string
	Price                       float32
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p *Product) GetAll() {
	fmt.Println("[")
	for _, product := range Products {
		fmt.Printf("\t %v \n", product)
	}
	fmt.Println("]")
}

func getByID(id int) (*Product, error) {
	for _, product := range Products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("ID %d not found", id)
}
