package models

import (
	"context"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Category struct {
	CategoryId int    `json:"accountId"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	Icon       string `json:"icon"`
	Type       string `json:"type"`
}

type CategoryModel struct {
	DB *database.DB
}

func (m *CategoryModel) Insert() (int, error) {
	// TODO
	return 0, nil
}

func (m *CategoryModel) Get(id int) (Category, error) {
	ctx := context.Background()
	stmt := `SELECT category_id, name, enabled, icon, type FROM category WHERE category_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var opt Category
	err := row.Scan(&opt.CategoryId, &opt.Name, &opt.Enabled, &opt.Icon, &opt.Type)
	if err != nil {
		return Category{}, err
	}

	return opt, nil
}
