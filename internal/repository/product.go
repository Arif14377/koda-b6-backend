package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ProductRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewProductRepository(db *pgxpool.Pool, rdb *redis.Client) *ProductRepository {
	return &ProductRepository{
		db:    db,
		redis: rdb,
	}
}

func (p *ProductRepository) GetAllProducts() (*[]models.Products, error) {
	ctx := context.Background()
	cacheKey := "products:all"

	// Try to get from cache
	cachedData, err := p.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var products []models.Products
		if err := json.Unmarshal([]byte(cachedData), &products); err == nil {
			return &products, nil
		}
	}

	query := `
		SELECT 
			p.id, p.name, p.description, p.quantity, p.price, p.rating, p.old_price, p.is_flash_sale,
			COALESCE((SELECT path FROM product_images WHERE product_id = p.id LIMIT 1), '') as image,
			COALESCE((
				SELECT string_agg(c.name, ',') 
				FROM categories c 
				JOIN product_category pc ON pc.category_id = c.id 
				WHERE pc.product_id = p.id
			), '') as categories_list,
			COALESCE((
				SELECT json_agg(json_build_object('id', id, 'name', name, 'addPrice', add_price))
				FROM product_variant WHERE product_id = p.id
			), '[]') as variants,
			COALESCE((
				SELECT json_agg(json_build_object('id', id, 'name', name, 'addPrice', add_price))
				FROM product_size WHERE product_id = p.id
			), '[]') as sizes
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
			&product.Image, &categoriesList, &product.Variants, &product.Sizes,
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

	// Save to cache
	if jsonData, err := json.Marshal(products); err == nil {
		p.redis.Set(ctx, cacheKey, jsonData, 1*time.Hour)
	}

	return &products, nil
}

func (p *ProductRepository) GetProductById(id int) (*models.Products, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("product:%d", id)

	// Try to get from cache
	cachedData, err := p.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var product models.Products
		if err := json.Unmarshal([]byte(cachedData), &product); err == nil {
			return &product, nil
		}
	}

	queryProduct := `
		SELECT id, name, description, quantity, price, rating, old_price, is_flash_sale
		FROM products
		WHERE id = $1
	`
	rowProduct := p.db.QueryRow(ctx, queryProduct, id)
	var product models.Products
	err = rowProduct.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.Rating, &product.OldPrice, &product.IsFlashSale)
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

	// Fetch Variants
	queryVariants := `SELECT id, name, add_price FROM product_variant WHERE product_id = $1`
	rowsVariants, err := p.db.Query(context.Background(), queryVariants, id)
	if err == nil {
		variants, err := pgx.CollectRows(rowsVariants, pgx.RowToStructByNameLax[models.ProductVariant])
		if err == nil {
			product.Variants = variants
		}
		rowsVariants.Close()
	}

	// Fetch Sizes
	querySizes := `SELECT id, name, add_price FROM product_size WHERE product_id = $1`
	rowsSizes, err := p.db.Query(context.Background(), querySizes, id)
	if err == nil {
		sizes, err := pgx.CollectRows(rowsSizes, pgx.RowToStructByNameLax[models.ProductSize])
		if err == nil {
			product.Sizes = sizes
		}
		rowsSizes.Close()
	}

	// Fetch Categories
	queryCategories := `
		SELECT c.name 
		FROM categories c 
		JOIN product_category pc ON pc.category_id = c.id 
		WHERE pc.product_id = $1
	`
	rowsCategories, err := p.db.Query(context.Background(), queryCategories, id)
	if err == nil {
		var categories []string
		for rowsCategories.Next() {
			var catName string
			if err := rowsCategories.Scan(&catName); err == nil {
				categories = append(categories, catName)
			}
		}
		product.Category = categories
		rowsCategories.Close()
	}

	// Save to cache
	if jsonData, err := json.Marshal(product); err == nil {
		p.redis.Set(ctx, cacheKey, jsonData, 1*time.Hour)
	}

	return &product, nil
}

func (p *ProductRepository) UpdateProduct(id int, product models.Products) error {
	ctx := context.Background()
	query := `
		UPDATE products 
		SET name=$1, description=$2, quantity=$3, price=$4, rating=$5, old_price=$6, is_flash_sale=$7
		WHERE id=$8
	`
	_, err := p.db.Exec(ctx, query,
		product.Name, product.Description, product.Quantity, product.Price,
		product.Rating, product.OldPrice, product.IsFlashSale, id,
	)
	if err != nil {
		return err
	}

	// Invalidate Cache
	p.redis.Del(ctx, "products:all")
	p.redis.Del(ctx, fmt.Sprintf("product:%d", id))

	return nil
}

func (p *ProductRepository) DeleteProduct(id int) error {
	ctx := context.Background()
	_, err := p.db.Exec(ctx, "DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return err
	}

	// Invalidate Cache
	p.redis.Del(ctx, "products:all")
	p.redis.Del(ctx, fmt.Sprintf("product:%d", id))

	return nil
}
