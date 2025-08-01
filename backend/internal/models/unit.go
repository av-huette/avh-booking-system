package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Unit struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UnitModel struct {
	DB *database.DB
}

func (m *UnitModel) Insert(name string) error {
	ctx := context.Background()
	query := `
        INSERT INTO unit (name) 
        VALUES ($1)`
	_, err := m.DB.Exec(ctx, query, name)

	return err
}

func (m *UnitModel) Get(id int) (*Unit, error) {
	ctx := context.Background()
	stmt := `SELECT unit_id, name FROM unit WHERE unit_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var unit Unit
	err := row.Scan(&unit.Id, &unit.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &unit, nil
}
