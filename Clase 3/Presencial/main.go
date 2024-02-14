package main

import "fmt"

/*
Employee
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un
pequeño programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composición con la estructura Person.
Realizar un método a la estructura Employee que se llame PrintEmployee(), lo que hará es realizar
la impresión de los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y,
por último, ejecutar el método PrintEmployee().
Si logramos realizar este pequeño programa, pudimos ayudar a la empresa a solucionar la gestión de
los empleados.
*/

type Person struct {
	ID                int
	Name, DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Employee %v", e)
}

func main() {
	person1 := Person{
		ID:          1,
		Name:        "Juan",
		DateOfBirth: "2006-01-02",
	}

	employee1 := Employee{
		ID:       1,
		Position: "Developer",
		Person:   person1,
	}
	employee1.PrintEmployee()

}
