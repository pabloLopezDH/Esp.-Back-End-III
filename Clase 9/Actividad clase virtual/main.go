package main

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez
instalada, utilizamos el siguiente comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"

go mod init actividad/virtual

*/
import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//1) Captura la solicitud GET “/hola-mundo”
	router.GET("/hola-mundo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "¡Hola Mundo!",
		})
	})

	//2) Captura la solicitud GET "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")

	//Ejecutar con: go run main.go
	//Para parar el servidor hacer: Ctrl + C
}
