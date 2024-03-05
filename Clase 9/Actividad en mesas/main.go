package main

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez
instalada, utilizamos el siguiente comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

// go mod init actividad/mesas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id              int
	Nombre          string
	Precio          float64
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func main() {
	router := gin.Default()

	//Captura la solicitud GET “/productos"
	router.GET("/productos", func(c *gin.Context) {

		datosBytes, err := ioutil.ReadFile("./productos.json")
		if err != nil {
			log.Fatal(err)
		}
		datosString := string(datosBytes)

		//Imprimo la lista de personas por consola
		fmt.Println(datosString)

		var listProducts []Productos

		if err := json.Unmarshal([]byte(datosString), &listProducts); err != nil {
			log.Fatal(err)
		}

		//Devolver Lista de Productos
		c.JSON(200, listProducts)
	})

	router.Run(":8080")

	//Ejecutar con: go run main.go
	//Para para el servidor hacer: Ctrl + C
}
