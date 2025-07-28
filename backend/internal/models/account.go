package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int              `json:"id"`
	FirstName string           `json:"firstName"`
	Nickname  string           `json:"nickname"`
	LastName  string           `json:"lastName"`
	Email     string           `json:"email"`
	Phone     string           `json:"phone"`
	Balance   pgtype.Numeric   `json:"balance"`
	MaxDebt   int              `json:"maxDebt"`
	Category  int              `json:"category"`
	Enabled   bool             `json:"enabled"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
}

type AccountModel struct {
	DB *database.DB
}

func (m *AccountModel) Insert(name string) error {
	ctx := context.Background()
	query := `
        INSERT INTO accounts (first_name, nickname, last_name, email, phone, balance, max_debt, category, enabled) 
        VALUES ('', $1, '', '', '', 0, 100, 1, true)`
	_, err := m.DB.Exec(ctx, query, name)

	return err
}

func (m *AccountModel) Get(id int) (*Account, error) {
	ctx := context.Background()
	stmt := `SELECT account_id, first_name, nickname, last_name, email, phone, balance,
       max_debt, category, enabled, created_at FROM accounts WHERE account_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var account Account
	err := row.Scan(&account.ID, &account.FirstName, &account.Nickname, &account.LastName, &account.Email,
		&account.Phone, &account.Balance, &account.MaxDebt, &account.Category, &account.Enabled, &account.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &account, nil
}
