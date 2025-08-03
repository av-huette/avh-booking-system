package models

import (
	"context"
	"fmt"
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
	"time"
)

var dbPool *database.DB
var accountModel *AccountModel

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	currentWorkDirectory, _ := os.Getwd()
	dbConf := config.LoadDbConfigFromFileEnv(string(currentWorkDirectory) + `/.env`)

	dbPool, err = database.New(dbConf.DbUser, dbConf.DbPassword, dbConf.DbHost, dbConf.DbPort, dbConf.DbName)
	if err != nil {
		panic(err)
	}

	setUp()

	accountModel = &AccountModel{DB: dbPool}

	defer tearDown()

	return m.Run(), nil
}

func getQueriesFromFile(filePath string) []string {
	// open file and read into string
	buf, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	s := string(buf)

	// separate queries
	s = strings.ReplaceAll(s, "\n", "")
	queries := strings.Split(s, ";")
	if queries[len(queries)-1] == "" {
		queries = queries[:len(queries)-1]
	}
	return queries
}

func batchExecQueries(queries []string) pgx.BatchResults {
	// batch queries
	batch := &pgx.Batch{}
	for _, query := range queries {
		batch.Queue(query)
	}

	// execute queries
	ctx := context.Background()
	br := dbPool.SendBatch(ctx, batch)
	_, err := br.Exec()
	if err != nil {
		panic(err)
	}

	return br
}

func setUp() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/create_tables.sql`
	queries := getQueriesFromFile(filePath)
	batchExecQueries(queries)
	logSetup("Created tables")

	filePath = currentWorkDirectory + `/insert_test_data.sql`
	queries = getQueriesFromFile(filePath)
	batchExecQueries(queries)
	logSetup("Inserted test data")
}

func tearDown() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/drop_tables.sql`
	queries := getQueriesFromFile(filePath)

	br := batchExecQueries(queries)

	err := br.Close()
	if err != nil {
		panic(err)
	}

	logTearDown("Dropped tables")

	// TODO: this is a workaround to close the dbPool as it is stuck in an endless loop
	done := make(chan struct{})
	go func() {
		dbPool.Close()
		close(done)
	}()

	select {
	case <-done:
		logTearDown("Database closed")
	case <-time.After(5 * time.Second):
		logTearDown("Database close timed out")
	}
}

func logSetup(msg string) {
	fmt.Println(fmt.Sprintf("== SETUP: %s", msg))
}

func logTearDown(msg string) {
	fmt.Println(fmt.Sprintf("== TEARDOWN: %s", msg))
}

func TestInsertAccount(t *testing.T) {
	dummyAccount := CreateAccount("FirstName", "NickName", "LastName",
		"Email", "Phone", "12.34", 100, 1)
	id, err := accountModel.Insert(dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccountById(t *testing.T) {
	_, err := accountModel.Get(1)
	// TODO verify fields
	require.NoError(t, err)
}
