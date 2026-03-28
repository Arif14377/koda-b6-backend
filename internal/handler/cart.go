package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Printf("Empty user ID in context")
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized access.",
		})
		return
	}

	cart, err := h.service.GetCartByUserId(c.Request.Context(), userId)
	if err != nil {
		log.Printf("Failed to get cart: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success get cart",
		Results: cart,
	})
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		log.Printf("Empty user ID in context")
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized access.",
		})
		return
	}

	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input data.",
		})
		return
	}

	cart.UserId = userId
	if err := h.service.AddToCart(c.Request.Context(), cart); err != nil {
		log.Printf("Failed to add to cart: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Success add to cart",
	})
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized access.",
		})
		return
	}

	cartId, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input data.",
		})
		return
	}

	if err := h.service.UpdateQuantity(c.Request.Context(), cartId, userId, input.Quantity); err != nil {
		log.Printf("Failed to update cart quantity: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success update quantity",
	})
}

func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized access.",
		})
		return
	}

	cartId, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.RemoveFromCart(c.Request.Context(), cartId, userId); err != nil {
		log.Printf("Failed to remove from cart: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success remove from cart",
	})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized access.",
		})
		return
	}

	if err := h.service.ClearCart(c.Request.Context(), userId); err != nil {
		log.Printf("Failed to clear cart: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success clear cart",
	})
}
