package models

import (
	"context"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	ProductId      int              `json:"productId"`
	Name           string           `json:"name"`
	Price          pgtype.Numeric   `json:"price"`
	ProductGroupId int              `json:"productGroupId"`
	Size           int              `json:"size"`
	UnitId         int              `json:"unitId"`
	Tax            pgtype.Numeric   `json:"tax"`
	CategoryId     int              `json:"categoryId"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
}

type ProductModel struct {
	DB *database.DB
}

func (m *ProductModel) Insert() (int, error) {
	// TODO
	return 0, nil
}

func (m *ProductModel) Get(productId int) (Product, error) {
	ctx := context.Background()
	stmt := `SELECT product_id, name, price, product_group, size, unit, tax, category, created_at FROM product WHERE product_id = $1`
	row := m.DB.QueryRow(ctx, stmt, productId)

	var product Product
	err := row.Scan(&product.ProductId, &product.Name, &product.Price, &product.ProductGroupId,
		&product.Size, &product.UnitId, &product.Tax, &product.CategoryId, &product.CreatedAt)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}
