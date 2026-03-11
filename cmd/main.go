package main

import (
	"log"

	"github.com/arif14377/koda-b6-backend/internal/di"
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

	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAllUser)
		users.POST("/by-email", userHandler.GetUserByEmail)
		users.POST("/forgot-password", forgotPasswordHandler.GenerateOTP)
	}

	r.Run("localhost:8888")
	// // AUTH
	// // register
	// r.POST("/register", handler.Register)
	// // login
	// r.POST("/login", handler.Login)
	// // get all users
	// r.GET("/users", handler.GetUsers)
	// // check user details
	// r.GET("/users/:id", handler.UserDetails)
	// // delete user
	// r.DELETE("/users/:id", handler.DeleteUser)
	// // update data user
	// r.PUT("/profile", handler.UpdateUser)

	// // PRODUCT
	// // get all products
	// r.GET("/products", handler.GetProducts)
	// // get product details
	// r.GET("/products/:id", handler.ProductDetails)
	// // add product
	// r.POST("/products", handler.AddProduct)
	// // delete product
	// r.DELETE("/products/:id", handler.DeleteProduct)
	// r.PUT("/products/:id", handler.UpdateProduct)

	// if err := handler.InitDB(); err != nil {
	// 	log.Fatalf("gagal connect db: %v", err)
	// }

	// // database := os.Getenv("DATABASE")
	// port := os.Getenv("PORT")
	// fmt.Println(port)

	// if err := r.Run(fmt.Sprintf("localhost:%s", port)); err != nil {
	// 	log.Fatalf("failed to run: %v", err)
	// }
}
