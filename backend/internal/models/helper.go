package models

import "github.com/jackc/pgx/v5/pgtype"

func NewNumeric(value string) pgtype.Numeric {
	var num pgtype.Numeric
	err := num.Scan(value)
	if err != nil {
		return pgtype.Numeric{}
	}
	return num
}
