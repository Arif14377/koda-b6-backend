package main

import (
	"log"

	"github.com/arif14377/koda-b6-backend1/internal/cors"
	"github.com/arif14377/koda-b6-backend1/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Middleware())

	// AUTH
	// register
	r.POST("/register", handler.Register)
	// login
	r.POST("/login", handler.Login)
	// get all users
	r.GET("/users", handler.GetUsers)
	// check user details
	r.GET("users/:id", handler.UserDetails)
	// delete user
	r.DELETE("/users/:id", handler.DeleteUser)
	// update data user
	r.PUT("/profile", handler.UpdateUser)

	// PRODUCT
	// get all products
	r.GET("/products", handler.GetProducts)
	// get product details
	r.GET("/products/:id", handler.ProductDetails)
	// add product
	r.POST("/products", handler.AddProduct)
	// delete product
	r.DELETE("/products/:id", handler.DeleteProduct)
	r.PUT("/products/:id", handler.UpdateProduct)
	if err := r.Run("localhost:8888"); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
