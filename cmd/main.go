package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arif14377/koda-b6-backend/internal/di"
	"github.com/arif14377/koda-b6-backend/internal/middleware"
	"github.com/arif14377/koda-b6-backend/internal/middleware/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	r.Use(cors.Middleware())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	container := di.NewCointainer()

	userHandler := container.UserHandler()
	forgotPasswordHandler := container.ForgotPasswordHandler()
	authHandler := container.AuthHandler()
	productHandler := container.ProductHandler()
	reviewHandler := container.ReviewHandler()
	cartHandler := container.CartHandler()
	transactionHandler := container.TransactionHandler()

	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("", userHandler.GetAllUser)
		users.POST("/by-email", userHandler.GetUserByEmail)
		users.POST("/forgot-password", forgotPasswordHandler.GenerateOTP)
		users.POST("/forgot-password/verifikasi-otp", forgotPasswordHandler.VerifikasiOTP)
		users.PATCH("/forgot-password/change", forgotPasswordHandler.ChangePassword)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	public := r.Group("/")
	{
		public.GET("/products", productHandler.GetAllProducts)
		public.GET("/products/:id", productHandler.GetProductById)
		public.GET("/reviews", reviewHandler.GetAllReviews)
	}

	cart := r.Group("/cart")
	cart.Use(middleware.AuthMiddleware())
	{
		cart.GET("", cartHandler.GetCart)
		cart.POST("", cartHandler.AddToCart)
		cart.PATCH("/:id", cartHandler.UpdateQuantity)
		cart.DELETE("/:id", cartHandler.RemoveFromCart)
		cart.DELETE("", cartHandler.ClearCart)
	}

	history := r.Group("/history")
	history.Use(middleware.AuthMiddleware())
	{
		history.GET("", transactionHandler.GetHistory)
		history.GET("/:id", transactionHandler.GetDetail)
		history.POST("", transactionHandler.Checkout)
	}

	r.GET("/delivery-methods", transactionHandler.GetDeliveryMethods)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Printf("%v", err)
		return
	}
}
