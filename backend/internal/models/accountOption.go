package models

import (
	"context"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type AccountOption struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AccountOptionModel struct {
	DB *database.DB
}

func (m *AccountOptionModel) Insert() (int, error) {
	// TODO
	return 0, nil
}

func (m *AccountOptionModel) Get(accountId int, key string) (AccountOption, error) {
	ctx := context.Background()
	stmt := `SELECT account, key, value FROM account_option WHERE account = $1 AND key = $2`
	row := m.DB.QueryRow(ctx, stmt, accountId, key)

	var opt AccountOption
	err := row.Scan(&opt.Id, &opt.Key, &opt.Value)
	if err != nil {
		return AccountOption{}, err
	}

	return opt, nil
}
