package main

import (
	"fmt"
	"time"
)

/*
Consigna
Una empresa de finanzas procesa órdenes de préstamos y órdenes de devoluciones
en una proporción de dos a uno (es decir, por cada devolución dan dos préstamos).
Aproximadamente, se procesa una devolución por segundo, y nos han solicitado
que programemos un modelo de procesamiento concurrente que denote este modelo.
Se requiere que corran al menos dos goroutines de forma concurrente que procesen
estas órdenes. Donde además se muestre por la consola una devolución por cada
dos préstamos hasta que se apriete cualquier botón y se acabe con la ejecución
del programa. El output esperado debería ser similar a este:
devolucion procesada
prestamo procesado
prestamo procesado
devolucion procesada
prestamo procesado
prestamo procesado
*/

func main() {

	devolucion := make(chan string)
	prestamo := make(chan string)

	go func() {
		defer close(devolucion)
		for {
			time.Sleep(time.Second)
			//fmt.Println("devolucion procesada")
			devolucion <- "devolucion procesada"
		}
	}()

	go func() {
		defer close(prestamo)
		for {
			time.Sleep(time.Second / 2)
			//fmt.Println("prestamo procesado")
			prestamo <- "prestamo procesado"
		}
	}()

	//fmt.Println(<-devolucion)
	//fmt.Println(<-prestamo)

	/*go func() {
		for devPrecesada := range devolucion {
			fmt.Println(devPrecesada)
		}
	}()

	go func() {
		for preProcesada := range prestamo {
			fmt.Println(preProcesada)
		}
	}()*/

	go func() {
		for {
			select {
			case msg := <-devolucion:
				fmt.Println(msg)
			case msg := <-prestamo:
				fmt.Println(msg)
			}
		}
	}()

	/*
		go func() {
			for {
				select {
				case msg := <-devolucion1:
					fmt.Println(msg)
				case msg := <-devolucion2:
					fmt.Println(msg)
				case msg := <-devolucion3:
					fmt.Println(msg)
				case err := <-errores:
					fmt.Println(err)
				}
			}
		}()
	*/

	fmt.Scanln()
}
