package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arif14377/koda-b6-backend/internal/models"
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
		log.Printf("Gagal mendapatkan semua produk: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List all products",
		Results: products,
	})
}

func (p *ProductHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	var productId int
	_, err := fmt.Sscanf(id, "%d", &productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "ID Produk tidak valid.",
		})
		return
	}

	product, err := p.productService.GetProductById(productId)
	if err != nil {
		log.Printf("Gagal mendapatkan detail produk: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Detail product",
		Results: product,
	})
}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var id int
	fmt.Sscanf(idStr, "%d", &id)

	var data models.Products
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input tidak valid.",
		})
		return
	}

	if err := p.productService.UpdateProduct(id, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal update produk.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Produk berhasil diupdate dan cache dibersihkan.",
	})
}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var id int
	fmt.Sscanf(idStr, "%d", &id)

	if err := p.productService.DeleteProduct(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal hapus produk.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Produk berhasil dihapus dan cache dibersihkan.",
	})
}
