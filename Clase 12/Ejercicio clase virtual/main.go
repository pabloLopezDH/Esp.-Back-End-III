package main

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez
instalada, utilizamos el siguiente comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

// go mod init actividad/virtual

// go mod tidy

//go get -u github.com/joho/godotenv

//import _ "github.com/joho/godotenv/autoload"

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")

	fmt.Print("My user: ", user, "\nMy password: ", pass)

	//Ejecutar con: go run main.go
	//Para terminar el servidor hacer: Ctrl + C
}
