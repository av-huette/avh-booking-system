package models

import (
	"context"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Categories struct {
	CategoryId int    `json:"accountId"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	Icon       string `json:"icon"`
}

type CategoriesModel struct {
	DB *database.DB
}

func (m *CategoriesModel) Insert() (int, error) {
	// TODO
	return 0, nil
}

func (m *CategoriesModel) Get(id int) (Categories, error) {
	ctx := context.Background()
	stmt := `SELECT category_id, name, enabled, icon FROM categories WHERE category_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var opt Categories
	err := row.Scan(&opt.CategoryId, &opt.Name, &opt.Enabled, &opt.Icon)
	if err != nil {
		return Categories{}, err
	}

	return opt, nil
}
