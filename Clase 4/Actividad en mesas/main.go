package main

import "fmt"

/*
Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar
productos y retornar el valor del precio total. La empresa tiene tres tipos de productos:
Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
Pequeño: solo tiene el costo del producto.
Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio, y retorne
una interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total
basándose en el costo del producto y los adicionales, en caso de que los tenga.
*/

type tiendaModa struct {
	nombre string
	precio float64
	tienda string
	url    string
}

type tiendacolumbia struct {
	tienda string
	precio float64
}

type Ecommerce interface {
	Precio(size string) float64
	Envio(direccion string) string
}

func main() {
	fmt.Println("------------------------------------------------")
	tiendaModa := nuevaTienda("tiendamoda")
	precioModa := tiendaModa.Precio("grande")
	envioModa := tiendaModa.Envio("Avenida Independencia 578")
	fmt.Println(tiendaModa)
	fmt.Println(envioModa)
	fmt.Println("Precio del producto:", precioModa)
	fmt.Println("------------------------------------------------")
	tiendaColumbia := nuevaTienda("tiendacolumbia")
	precioColumbia := tiendaColumbia.Precio("Mediano")
	envioColumbia := tiendaColumbia.Envio("Avenida Independencia 578")
	fmt.Println(tiendaColumbia)
	fmt.Println(envioColumbia)
	fmt.Println("Precio del pruducto:", precioColumbia)
	fmt.Println("------------------------------------------------")
}

func (t tiendaModa) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return t.precio
	case "mediano":
		var porcentaje float64
		porcentaje = (t.precio / 100) * 3
		return t.precio + porcentaje
	case "grande":
		var porcentaje float64
		flete := 2500
		porcentaje = (t.precio / 100) * 6
		return t.precio + porcentaje + float64(flete)
	}
	return t.precio
}

func (t tiendaModa) Envio(dir string) string {
	enviado := "Enviando un paquete a " + dir
	return enviado
}

func (t tiendacolumbia) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return t.precio
	case "mediano":
		var porcentaje float64
		porcentaje = (t.precio / 100) * 3
		return t.precio + porcentaje
	case "grande":
		var porcentaje float64
		flete := 2500
		porcentaje = (t.precio / 100) * 6
		return t.precio + porcentaje + float64(flete)
	}
	return t.precio
}

func (t tiendacolumbia) Envio(dir string) string {
	enviado := "Enviando un paquete a " + dir
	return enviado
}

func nuevaTienda(mailType string) Ecommerce {
	if mailType == "tiendamoda" {
		return tiendaModa{nombre: "mesa", precio: 2000, tienda: "sucursal Colombia", url: "http//:tiendaModa.com"}
	}
	if mailType == "tiendacolumbia" {
		return tiendaModa{tienda: "sucursal Mexico", precio: 5000}
	}
	return nil
}
