package main

/*
Para utilizar Gin se requiere la versión 1.13+ de Go, una vez
instalada, utilizamos el siguiente comando para instalar Gin:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

// go mod init actividad/mesas

// go mod tidy

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var productsList = []Product{}

func main() {
	loadProducts("products.json", &productsList)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", GetAllProducts())
		products.GET(":id", GetProduct())
		products.GET("/search", SearchProduct())

		products.GET("/productparams", GetProductByParams())
		products.GET("/searchbyquantity", SearchProductByQuantity())
		products.GET("/buy", Buy())

	}
	r.Run(":8080")

	//Ejecutar con: go run main.go
	//Para para el servidor hacer: Ctrl + C
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

// GetAllProducts traer todos los productos almacenados
func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, productsList)
	}
}

// GetProduct traer un producto por id
func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		for _, product := range productsList {
			if product.Id == id {
				ctx.JSON(200, product)
				return
			}
		}
		ctx.JSON(404, gin.H{"error": "product not found"})
	}
}

// SearchProduct traer los productos mayores a un precio tomado por parametro
func SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid price"})
			return
		}
		list := []Product{}
		for _, product := range productsList {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}
		ctx.JSON(200, list)
	}
}

// GetProductByParams asigna a los campos de nuestra estructura con los 
//datos que recibimos de los params de nuestra request.
func GetProductByParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	var product Product

	id := ctx.Query("id")
	parseId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id."})
		return
	}
	product.Id = parseId

	name := ctx.Query("name")
	product.Name = name

	quantity := ctx.Query("quantity")
	parseQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Quantity."})
		return
	}
	product.Quantity = parseQuantity

	code_value := ctx.Query("code_value")
	product.CodeValue = code_value

	is_published := ctx.Query("is_published")
	parseIs_published, err := strconv.ParseBool(is_published)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid is_published."})
		return
	}
	product.IsPublished = parseIs_published

	expiration := ctx.Query("expiration")
	product.Expiration = expiration

	price := ctx.Query("price")
	parsePrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid price."})
		return
	}
	product.Price = parsePrice

	productsList = append(productsList, product)

	if ctx.BindQuery(&product) == nil {
		ctx.JSON(200, gin.H{
			"product": product,
		})
	}
	}
}

// SearchProductByQuantity traer la lista de los productos que esten 
//entre dos limites de cantidades.
func SearchProductByQuantity() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		from := ctx.Query("from")
		parseFrom, err := strconv.Atoi(from)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Quantity."})
			return
		}

		to := ctx.Query("to")
		parseTo, err := strconv.Atoi(to)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Quantity."})
			return
		}

		list := []Product{}
		for _, product := range productsList {
			if product.Quantity > parseFrom && product.Quantity < parseTo {
				list = append(list, product)
			}
		}
		ctx.JSON(200, list)
	}
}

// Buy traer el detalle de una compra al buscar por code_valu y el resultado total de la compra 
func Buy() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		code_value:= ctx.Query("code_value")

		quantity := ctx.Query("quantity")
		parseQuantity, err := strconv.ParseFloat(quantity, 32) 
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Quantity."})
			return
		}

		for _, product := range productsList {
			if product.CodeValue == code_value {
				var totalPrice = product.Price * parseQuantity 

				ctx.JSON(200, gin.H{
					"product": product.Name,
					"quantity": parseQuantity,
					"total": totalPrice,
				})
			}
		}	
	}
}