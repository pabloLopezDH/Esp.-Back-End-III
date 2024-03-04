package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

func main() {

	router := gin.Default()

	jsonPer := `{"Nombre": "Juan", "Apellido": "Perez", "Edad": 21, "Direccion": "calle falsa 123", "Telefono": "42554477", "Activo": true}`

	var p persona

	if err := json.Unmarshal([]byte(jsonPer), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	router.GET("/persona", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"persona": p,
		})
		//c.JSON(200, p)
	})

	router.GET("/persona2", func(c *gin.Context) {
		panic("test")
		c.JSON(200, gin.H{
			"persona": p,
		})
	})

	router.Run(":8080")

}
