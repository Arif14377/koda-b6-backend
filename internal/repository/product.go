package repository

import (
	"context"
	"strings"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetAllProducts() (*[]models.Products, error) {
	query := `
		SELECT 
			p.id, p.name, p.description, p.quantity, p.price, p.rating, p.old_price, p.is_flash_sale,
			COALESCE((SELECT path FROM product_images WHERE product_id = p.id LIMIT 1), '') as image,
			COALESCE((
				SELECT string_agg(c.name, ',') 
				FROM categories c 
				JOIN product_category pc ON pc.category_id = c.id 
				WHERE pc.product_id = p.id
			), '') as categories_list
		FROM products p
	`
	rows, err := p.db.Query(context.Background(), query)
	if err != nil {
		return &[]models.Products{}, err
	}
	defer rows.Close()

	var products []models.Products
	for rows.Next() {
		var product models.Products
		var categoriesList string
		err := rows.Scan(
			&product.Id, &product.Name, &product.Description, &product.Quantity,
			&product.Price, &product.Rating, &product.OldPrice, &product.IsFlashSale,
			&product.Image, &categoriesList,
		)
		if err != nil {
			return &[]models.Products{}, err
		}
		if categoriesList != "" {
			product.Category = strings.Split(categoriesList, ",")
		} else {
			product.Category = []string{}
		}
		// Promo logic (misal jika is_flash_sale true, tambahkan ke promo)
		product.Promo = []string{}
		if product.IsFlashSale {
			product.Promo = append(product.Promo, "Flash Sale")
		}
		if product.OldPrice > product.Price {
			product.Promo = append(product.Promo, "Cheap")
		}

		products = append(products, product)
	}

	return &products, nil
}

func (p *ProductRepository) GetProductById(id int) (*models.Products, error) {
	queryProduct := `
		SELECT id, name, description, quantity, price, rating, old_price, is_flash_sale
		FROM products
		WHERE id = $1
	`
	rowProduct := p.db.QueryRow(context.Background(), queryProduct, id)
	var product models.Products
	err := rowProduct.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.Rating, &product.OldPrice, &product.IsFlashSale)
	if err != nil {
		return nil, err
	}

	// Fetch Images
	queryImages := `SELECT id, product_id, path FROM product_images WHERE product_id = $1`
	rowsImages, err := p.db.Query(context.Background(), queryImages, id)
	if err == nil {
		images, err := pgx.CollectRows(rowsImages, pgx.RowToStructByNameLax[models.ProductImage])
		if err == nil {
			product.Images = images
			if len(images) > 0 {
				product.Image = images[0].Path
			}
		}
		rowsImages.Close()
	}

	// Fetch Variants (Assuming variants are shared or linked, but here we just show what's needed for the UI)
	// If variants are not in DB yet, we can mock them or leave empty.
	// Based on seed.sql (if exists), let's check what variants/sizes are available.

	return &product, nil
}
