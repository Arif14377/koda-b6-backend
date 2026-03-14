package handler

// import (
// 	"fmt"
// 	"slices"
// 	"strconv"
// 	"strings"

// 	"github.com/arif14377/koda-b6-backend/internal/models"
// 	"github.com/gin-gonic/gin"
// )

// var listProducts []models.Products

// // get all products
// func GetProducts(c *gin.Context) {
// 	c.JSON(200, models.Response{
// 		Success: true,
// 		Message: "List products",
// 		Results: listProducts,
// 	})
// }

// // get product details
// func ProductDetails(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	product := models.Products{}
// 	isExist := false

// 	for _, p := range listProducts {
// 		if p.Id == id {
// 			product = p
// 			isExist = true
// 			break
// 		}
// 	}

// 	if !isExist {
// 		c.JSON(404, models.Response{
// 			Success: false,
// 			Message: "Product tidak ditemukan.",
// 		})
// 		return
// 	}

// 	c.JSON(200, models.Response{
// 		Success: true,
// 		Message: fmt.Sprintf("Product detail id: %d", id),
// 		Results: product,
// 	})
// }

// // add product
// func AddProduct(c *gin.Context) {
// 	data := models.Products{}
// 	err := c.ShouldBindJSON(&data)

// 	if err != nil {
// 		c.JSON(401, models.Response{
// 			Success: false,
// 			Message: "JSON tidak valid",
// 		})
// 		return
// 	}

// 	for _, p := range listProducts {
// 		if strings.Contains(p.Name, data.Name) {
// 			c.JSON(400, models.Response{
// 				Success: false,
// 				Message: "Nama product sudah ada.",
// 			})
// 			return
// 		}
// 	}

// 	if data.Name == "" || data.Description == "" || data.Qty == 0 || data.Price == 0 {
// 		c.JSON(400, models.Response{
// 			Success: false,
// 			Message: "Data tidak boleh kosong.",
// 		})
// 		return
// 	}

// 	data.Id = len(listProducts) + 1
// 	listProducts = append(listProducts, data)
// 	c.JSON(200, models.Response{
// 		Success: true,
// 		Message: "Produk berhasil ditambahkan.",
// 	})
// }

// // delete product
// func DeleteProduct(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	isExist := false

// 	for i, p := range listProducts {
// 		if p.Id == id {
// 			listProducts = slices.Delete(listProducts, i, i+1)
// 			isExist = true
// 			break
// 		}
// 	}

// 	if !isExist {
// 		c.JSON(404, models.Response{
// 			Success: false,
// 			Message: "Id produk tidak ditemukan",
// 		})
// 		return
// 	}

// 	c.JSON(200, models.Response{
// 		Success: true,
// 		Message: fmt.Sprintf("Data id ke-%d berhasil dihapus.", id),
// 	})
// }

// // update product
// func UpdateProduct(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	data := models.Products{}
// 	err := c.ShouldBindJSON(&data)

// 	if err != nil {
// 		c.JSON(401, models.Response{
// 			Success: false,
// 			Message: "JSON tidak valid.",
// 		})
// 		return
// 	}

// 	for _, p := range listProducts {
// 		if strings.Contains(p.Name, data.Name) {
// 			c.JSON(400, models.Response{
// 				Success: false,
// 				Message: "Nama product sudah ada.",
// 			})
// 			return
// 		}
// 	}

// 	if data.Name == "" || data.Description == "" || data.Price == 0 {
// 		c.JSON(400, models.Response{
// 			Success: false,
// 			Message: "Data tidak boleh kosong.",
// 		})
// 		return
// 	}

// 	for i, p := range listProducts {
// 		if p.Id == id {
// 			listProducts[i] = data
// 			listProducts[i].Id = id
// 			break
// 		}
// 	}

// 	c.JSON(200, models.Response{
// 		Success: true,
// 		Message: "Data berhasil diupdate.",
// 	})
// }
