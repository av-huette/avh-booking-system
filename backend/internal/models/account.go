package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	Id        int              `json:"id"`
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

func CreateAccount(firstName string, nickName string, lastName string,
	email string, phone string, balance string, maxDebt int, category int) Account {
	return Account{
		FirstName: firstName,
		Nickname:  nickName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Balance:   NewNumeric(balance),
		MaxDebt:   maxDebt,
		Category:  category,
		Enabled:   true,
	}
}

type AccountModel struct {
	DB *database.DB
}

func (m *AccountModel) Insert(account Account) error {
	ctx := context.Background()
	query := `
        INSERT INTO account (first_name, nickname, last_name, email, phone, balance, max_debt, category, enabled) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := m.DB.Exec(ctx, query, account.FirstName,
		account.Nickname,
		account.LastName,
		account.Email,
		account.Phone,
		account.Balance,
		account.MaxDebt,
		account.Category,
		account.Enabled)

	return err
}

func (m *AccountModel) Get(accountId int) (*Account, error) {
	ctx := context.Background()
	stmt := `SELECT account_id, first_name, nickname, last_name, email, phone, balance,
       max_debt, category, enabled, created_at FROM account WHERE account_id = $1`
	row := m.DB.QueryRow(ctx, stmt, accountId)

	var account Account
	err := row.Scan(&account.Id, &account.FirstName, &account.Nickname, &account.LastName, &account.Email,
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
