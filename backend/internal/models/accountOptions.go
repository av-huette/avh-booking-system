package models

import (
	"context"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type AccountOptions struct {
	AccountId int    `json:"accountId"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type AccountOptionsModel struct {
	DB *database.DB
}

func (m *AccountOptionsModel) Insert() (int, error) {
	// TODO
	return 0, nil
}

func (m *AccountOptionsModel) Get(accountId int, key string) (AccountOptions, error) {
	ctx := context.Background()
	stmt := `SELECT account, key, value FROM account_options WHERE account = $1 AND key = $2`
	row := m.DB.QueryRow(ctx, stmt, accountId, key)

	var opt AccountOptions
	err := row.Scan(&opt.AccountId, &opt.Key, &opt.Value)
	if err != nil {
		return AccountOptions{}, err
	}

	return opt, nil
}
