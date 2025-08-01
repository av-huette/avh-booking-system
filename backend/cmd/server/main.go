package main

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/logger"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log/slog"
	"os"
)

type dbModels struct {
	account        *models.AccountModel
	accountOptions *models.AccountOptionsModel
	categories     *models.CategoriesModel
}

type application struct {
	conf     *config.AppConfig
	log      *slog.Logger
	db       *database.DB
	dbModels dbModels
}

func main() {
	dbPool, err := database.New()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	app := &application{
		conf: config.LoadConfig(),
		log:  logger.CreateLogger(),
		db:   dbPool,
		dbModels: dbModels{
			&models.AccountModel{DB: dbPool},
			&models.AccountOptionsModel{DB: dbPool},
			&models.CategoriesModel{DB: dbPool}},
	}

	err = app.dbModels.account.Insert("Dummy")
	if err != nil {
		panic(err.Error())
	}

	categories, err := app.dbModels.categories.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("category: %#v", categories))

	account, err := app.dbModels.account.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("account: %#v", account))

	accountOptions, err := app.dbModels.accountOptions.Get(1, "opt1")
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("accountOption: %#v", accountOptions))

	err = app.serveHTTP()
	if err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
