package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type ProductGroup struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
}

type ProductGroupModel struct {
	DB *database.DB
}

func CreateProductGroup(name string, parentId int) ProductGroup {
	return ProductGroup{
		Name:     name,
		ParentId: parentId,
	}
}

func (m *ProductGroupModel) Insert(group ProductGroup) (int, error) {
	ctx := context.Background()
	var id int
	var err error
	if group.ParentId <= 0 {
		query := `
        INSERT INTO product_group (name) 
        VALUES ($1)
        RETURNING product_group_id`
		err = m.DB.QueryRow(ctx, query, group.Name).Scan(&id)
	} else {
		query := `
        INSERT INTO product_group (name, parent) 
        VALUES ($1, $2)
        RETURNING product_group_id`
		err = m.DB.QueryRow(ctx, query, group.Name, group.ParentId).Scan(&id)
	}

	return id, err
}

func (m *ProductGroupModel) Get(id int) (*ProductGroup, error) {
	ctx := context.Background()
	stmt := `SELECT product_group_id, name, parent
			FROM product_group
			WHERE product_group_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var productGroup ProductGroup
	var parentId *int // pointer to read null
	err := row.Scan(&productGroup.Id, &productGroup.Name, &parentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	if parentId == nil {
		productGroup.ParentId = 0
	} else {
		productGroup.ParentId = *parentId
	}

	return &productGroup, nil
}
