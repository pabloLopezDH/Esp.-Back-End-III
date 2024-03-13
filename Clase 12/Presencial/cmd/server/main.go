package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bootcamp-go/Consignas-Go-Web.git/cmd/server/handler"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}

	//os.Setenv("TOKEN", "my-secret-key")
	//fmt.Println(os.Getenv("TOKEN"))

	var productsList = []domain.Product{}
	loadProducts("../../products.json", &productsList)

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.PUT(":id", productHandler.Put())
		products.PATCH(":id", productHandler.Patch())
		products.DELETE(":id", productHandler.Delete())
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}
