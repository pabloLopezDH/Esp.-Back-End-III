package main

import (
	"fmt"
)

/*
Impuestos de salario
Repetir el proceso de la ejercitación realizada en clase, pero ahora implementando fmt.Errorf()
para que el mensaje de error reciba por parámetro el valor de salary indicando que no alcanza el mínimo imponible.
El mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y
el salario ingresado es de: [salary]” (siendo [salary] el valor de tipo int pasado por parámetro).
*/

var ErrSalary = fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de")

func validateSalary(salary int) error {
	if salary <= 150000 {
		return ErrSalary
	}
	return nil
}

func main() {

	salary := 100
	err := validateSalary(salary)
	if err != nil {
		fmt.Printf("%v: %v ", ErrSalary, salary)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
