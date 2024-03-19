package main

import (
	"github.com/bootcamp-go/Consignas-Go-Web.git/cmd/server/handler"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/middleware"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./cmd/server/.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	storage := store.NewStore("./products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.GET("/consumer_price", productHandler.ConsumerPrice())
		products.POST("", middleware.Authentication(), productHandler.Post())
		products.DELETE(":id", middleware.Authentication(), productHandler.Delete())
		products.PATCH(":id", middleware.Authentication(), productHandler.Patch())
		products.PUT(":id", middleware.Authentication(), productHandler.Put())
	}
	r.Run(":8080")
}
