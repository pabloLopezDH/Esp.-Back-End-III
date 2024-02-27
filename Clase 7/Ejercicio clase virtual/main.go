package main

import "fmt"

/*
  Calcular Precio
  Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
  Para ello requieren realizar un programa que se encargue de calcular el precio total de
  Productos, Servicios y Mantenimientos.
  Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la
  sumatoria se realice en paralelo mediante 3 go routines.

  Se requieren 3 estructuras:
  Productos: nombre, precio, cantidad.
  Servicios: nombre, precio, minutos trabajados.
  Mantenimiento: nombre, precio.

  Se requieren 3 funciones:
  Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
  Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
  si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
  Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

  Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/

func main() {
	outProducto := make(chan float32)
	outServicio := make(chan float32)
	outMantenimiento := make(chan float32)

	var total float32
	go sumarProductos([]Producto{
		{"Zapatilla", 4, 2},
		{"Remera", 2, 4},
		{"Pantalon", 4, 1},
	}, outProducto)

	go sumarServicios([]Servicio{
		{"Electricidad", 10, 30},
		{"General", 5, 90},
		{"Jardineria", 15, 15},
	}, outServicio)

	go sumarMantenimientos([]Mantenimiento{
		{"Pintura", 20},
		{"Cerraduras", 15},
		{"Vidrios", 50},
	}, outMantenimiento)

	for msg := range outProducto {
		total += msg
	}

	for msg := range outServicio {
		total += msg
	}

	for msg := range outMantenimiento {
		total += msg
	}

	fmt.Println(total)
}

type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad uint32
}

type Servicio struct {
	Nombre            string
	Precio            float32
	MinutosTrabajados int32
}

type Mantenimiento struct {
	Nombre string
	Precio float32
}

func sumarProductos(productos []Producto, out chan float32) {
	defer close(out)
	var precioTotal float32
	for _, p := range productos {
		precioTotal += p.Precio * float32(p.Cantidad)
	}
	out <- precioTotal
}

func sumarServicios(servicios []Servicio, out chan float32) {
	defer close(out)
	var precioTotal float32

	for _, s := range servicios {
		if s.MinutosTrabajados >= 30 {
			precioTotal += s.Precio * float32(s.MinutosTrabajados/30)
		} else {
			precioTotal += s.Precio
		}
	}
	out <- precioTotal
}

func sumarMantenimientos(mantenimientos []Mantenimiento, out chan float32) {
	defer close(out)
	var precioTotal float32

	for _, m := range mantenimientos {
		precioTotal += m.Precio
	}
	out <- precioTotal
}
