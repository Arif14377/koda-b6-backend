package service

import (
	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (p *ProductService) GetAllProducts() (*[]models.Products, error) {
	products, err := p.productRepo.GetAllProducts()
	if err != nil {
		return &[]models.Products{}, err
	}

	return products, nil
}
