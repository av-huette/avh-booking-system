package test

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var dbModels *modelStructs

func TestMain(m *testing.M) {
	dbModels = &modelStructs{}
	code, err := run(m, dbModels)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

// --------------------------------------------------
// Account
// --------------------------------------------------

func TestInsertAccount(t *testing.T) {
	dummyAccount := models.CreateAccount("一嫂", "Lady Hurricane", "鄭",
		"zheng.yi.sao@redfleet.cn", "+86 20 1807-2010", "12.34", 100, 1)
	id, err := dbModels.account.Insert(dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccountById(t *testing.T) {
	const accountId = 1
	daGama, err := dbModels.account.Get(accountId)
	if daGama == nil {
		t.Fail()
		t.Log("Could not get account")

		return
	}

	require.NoError(t, err)
	require.Equal(t, daGama.Id, accountId)
	require.Equal(t, daGama.FirstName, "Vasco")
	require.Equal(t, daGama.Nickname, "Cape Conquerer")
	require.Equal(t, daGama.LastName, "da Gama")
	require.Equal(t, daGama.Email, "indianspice@capeofgoodhope.com")
	require.Equal(t, daGama.Phone, "+351 914 97 1498")
	var expectedBalance pgtype.Numeric
	err = expectedBalance.Scan("33.55")
	if err != nil {
		t.Fatal("Could not convert string to Numeric")
	}
	require.Equal(t, daGama.Balance, expectedBalance)
	require.Equal(t, daGama.MaxDebt, 100)
	require.Equal(t, daGama.Category, 1)
	require.Equal(t, daGama.Enabled, true)
	require.NotZero(t, daGama.CreatedAt)
}

// --------------------------------------------------
// Category
// --------------------------------------------------

// TODO
