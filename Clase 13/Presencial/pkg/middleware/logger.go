package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
●	Verbo utilizado: GET, POST, PUT, etc.
●	Fecha y hora: podemos utilizar el paquete time.
●	URL de la consulta: localhost:8080/products
●	Tamaño en bytes: tamaño de la consulta.
*/
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		verb := c.Request.Method
		time1 := time.Now()
		path := c.Request.URL.Path
		c.Next()

		var size int
		if c.Writer != nil {
			size = c.Writer.Size()
		}
		time2 := time.Since(time1)

		fmt.Printf("\nverb: %v\ntime: %v\npath: %v\nsize: %v\ntime: %v\n", verb, time1, path, size, time2)
		c.Next()
	}
}
