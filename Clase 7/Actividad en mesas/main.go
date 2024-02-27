package main

import (
	"fmt"
)

/*
Consigna
Se requiere un programa para calcular si un número es par o impar desde una lista de int.
Para esto se deben crear dos funciones:
●	par(chan int) debe recibir a través de un channel de int los valores pares enviados y
mostrarlos por la terminal.
●	impar(chan int) debe recibir a través de un channel los valores impares y mostrarlos por terminal.
El programa debe comenzar dos goroutines que se ejecuten concurrentemente.
En la goroutine principal se itera el slice provisto con valores y, en caso de que el valor sea par,
se lo debe enviar al channel correspondiente. De igual manera para los impares.
Ayuda: Si el resto de una división por dos da cero es un número par (i%2 == 0) y si da 1 es impar
(i%2 == 0 ). Para obtener el resto de una operación, utilizar el operador %: i%2 == 1.
¡Éxitos!
*/

func main() {
	var intSlice = []int{12, 43, 34, 83, 54, 78, 90, 05, 100}
	chImpar := make(chan int)
	chPar := make(chan int)

	go impar(chImpar)
	go par(chPar)

	for _, value := range intSlice {
		if value%2 != 0 {
			chImpar <- value
		} else {
			chPar <- value
		}
	}

}

func impar(ch <-chan int) {
	for v := range ch {
		fmt.Println("Impar:", v)
	}
}

func par(ch <-chan int) {
	for v := range ch {
		fmt.Println("Par:", v)
	}
}
