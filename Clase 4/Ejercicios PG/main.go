package main

import (
	"fmt"
)

/*
Consigna
Una empresa de redes sociales requiere implementar una estructura usuarios con funciones
que vayan agregando información a la misma. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y
para las funciones. La estructura debe tener los campos: nombre, apellido, edad, correo
y contraseña. Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña.
*/

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func main() {
	fmt.Println("Agregar usuario: ")

	user := usuario{
		nombre:     "Juan",
		apellido:   "Peres",
		edad:       27,
		correo:     "juanperez@gmail.com",
		contraseña: "123facil",
	}
	fmt.Println(user)

	var usPuntero *usuario

	usPuntero = &user

	fmt.Println("\nCambiar nombre:")
	fmt.Println(usPuntero.cambiarNombre("Juan Martin", "Perez"))

	fmt.Println("Cambiar edad:")
	fmt.Println(usPuntero.cambiarEdad(28))

	fmt.Println("Cambiar correo:")
	fmt.Println(usPuntero.cambiarCorreo("jperez28@gmail.com"))

	fmt.Println("Cambiar contraseña:")
	fmt.Println(usPuntero.cambiarContraseña("456dificil"))

	fmt.Println("\nUser cambiado totalmente:")
	fmt.Println(user)

}

func (c *usuario) cambiarNombre(nombre, apellido string) string {
	c.nombre = nombre
	c.apellido = apellido
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarEdad(edad int) string {
	c.edad = edad
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarCorreo(correo string) string {
	c.correo = correo
	return fmt.Sprintln(c)
}

func (c *usuario) cambiarContraseña(contraseña string) string {
	c.contraseña = contraseña
	return fmt.Sprintln(c)
}
