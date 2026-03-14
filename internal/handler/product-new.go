package handler

import (
	"net/http"

	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (p *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":  false,
			"messages": "Failed to get all products.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"messages": "List all products",
		"results":  products,
	})
}
