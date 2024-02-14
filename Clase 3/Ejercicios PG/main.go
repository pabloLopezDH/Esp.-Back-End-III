package main

import (
	"fmt"
)

/*
Registro de estudiantes
Una universidad necesita registrar a los estudiantes y generar una funcionalidad para imprimir
el detalle de los datos de cada uno de ellos, de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos.
Para ello es necesario generar una estructura Alumno con las variables Nombre, Apellido, DNI, Fecha
y que tenga un método detalle.
*/

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (a Alumno) Detalle() {

	fmt.Printf("\nDetalle del alumno:")
	fmt.Printf("\nNombre: %s \nApellido: %s \nDNI: %v \nFecha: %s\n", a.nombre, a.apellido, a.dni, a.fecha)

}

func main() {

	alumno := Alumno{
		nombre:   "Carlos",
		apellido: "Molina",
		dni:      95866074,
		fecha:    "25/10/2023",
	}

	alumno2 := Alumno{
		nombre:   "Pablo",
		apellido: "Gonzalez",
		dni:      36564453,
		fecha:    "29/01/2015",
	}

	Alumno.Detalle(alumno)
	Alumno.Detalle(alumno2)

}
