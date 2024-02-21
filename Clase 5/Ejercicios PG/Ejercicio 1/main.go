package main

import "fmt"

/*
Ejercicio 1 - Impuestos de salario
En la función main, definir una variable salary y asignarle un valor de tipo “int”.
Luego, crear un error personalizado con un struct que implemente Error() con el mensaje
“Error: el salario ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary
sea menor a 150.000. De lo contrario, imprimir por consola el mensaje “Debe pagar impuesto”.
*/

type ErrSalary struct {
	message string
}

func (e *ErrSalary) Error() string {
	return e.message
}

func main() {
	salary := 100
	err := validarSalario(salary)
	if err != nil {
		fmt.Printf("%v", err.Error())
	} else {
		fmt.Println("Debe pagar impuestos.")
	}
}

func validarSalario(salary int) error {
	if salary <= 150000 {
		return &ErrSalary{message: "Error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return nil
}
