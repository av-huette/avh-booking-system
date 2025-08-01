package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type ProductGroup struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductGroupModel struct {
	DB *database.DB
}

func (m *ProductGroupModel) Insert(name string) error {
	ctx := context.Background()
	query := `
        INSERT INTO product_group (name) 
        VALUES ($1)`
	_, err := m.DB.Exec(ctx, query, name)

	return err
}

func (m *ProductGroupModel) Get(id int) (*ProductGroup, error) {
	ctx := context.Background()
	stmt := `SELECT product_group_id, name FROM product_group WHERE product_group_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var ProductGroup ProductGroup
	err := row.Scan(&ProductGroup.Id, &ProductGroup.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &ProductGroup, nil
}
