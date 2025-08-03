package test

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
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

func TestInsertAccount(t *testing.T) {
	dummyAccount := models.CreateAccount("FirstName", "NickName", "LastName",
		"Email", "Phone", "12.34", 100, 1)
	id, err := dbModels.account.Insert(dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccountById(t *testing.T) {
	_, err := dbModels.account.Get(1)
	// TODO verify fields
	require.NoError(t, err)
}
