package di

import (
	"context"
	"fmt"
	"os"

	"github.com/arif14377/koda-b6-backend/internal/handler"
	"github.com/arif14377/koda-b6-backend/internal/repository"
	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	db          *pgx.Conn
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handler.UserHandler

	fpRepo    *repository.ForgotPasswordRepository
	fpService *service.ForgotPasswordService
	fpHandler *handler.ForgotPasswordHandler

	aRepo    *repository.AuthRepository
	aService *service.AuthService
	aHandler *handler.AuthHandler

	pRepo    *repository.ProductRepository
	pService *service.ProductService
	pHandler *handler.ProductHandler
}

func NewCointainer() *Container {
	// fmt.Println(os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("Failed to connect database: %v", err)
		os.Exit(1)
	}

	container := Container{
		db: conn,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.db)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)

	c.fpRepo = repository.NewForgotPasswordRepository(c.db)
	c.fpService = service.NewForgotPasswordService(c.fpRepo, c.userRepo)
	c.fpHandler = handler.NewForgotPasswordHandler(c.fpService, c.userService)

	c.aRepo = repository.NewAuthRepository(c.db)
	c.aService = service.NewAuthService(c.aRepo, c.userRepo)
	c.aHandler = handler.NewAuthHandler(c.aService)

	c.pRepo = repository.NewProductRepository(c.db)
	c.pService = service.NewProductService(c.pRepo)
	c.pHandler = handler.NewProductHandler(c.pService)

}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) ForgotPasswordHandler() *handler.ForgotPasswordHandler {
	return c.fpHandler
}

func (c *Container) AuthHandler() *handler.AuthHandler {
	return c.aHandler
}

func (c *Container) ProductHandler() *handler.ProductHandler {
	return c.pHandler
}
