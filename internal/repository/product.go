package repository

import (
	"context"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetAllProducts() (*[]models.Products, error) {
	rows, err := p.db.Query(context.Background(), "SELECT id, name, description, quantity, price FROM products")
	// fmt.Println(rows)
	// fmt.Println(err)
	if err != nil {
		return &[]models.Products{}, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Products])
	// fmt.Println(products)
	if err != nil {
		return &[]models.Products{}, err
	}

	return &products, nil
}
