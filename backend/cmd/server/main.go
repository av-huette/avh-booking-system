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

type application struct {
	conf           *config.AppConfig
	log            *slog.Logger
	db             *database.DB
	account        *models.AccountModel
	accountOptions *models.AccountOptionsModel
	categories     *models.CategoriesModel
}

func main() {
	dbPool, err := database.New()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	app := &application{
		conf:           config.LoadConfig(),
		log:            logger.CreateLogger(),
		db:             dbPool,
		account:        &models.AccountModel{DB: dbPool},
		accountOptions: &models.AccountOptionsModel{DB: dbPool},
		categories:     &models.CategoriesModel{DB: dbPool},
	}

	err = app.account.Insert("Dummy")
	if err != nil {
		panic(err.Error())
	}

	categories, err := app.categories.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("category: %#v", categories))

	account, err := app.account.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("account: %#v", account))

	accountOptions, err := app.accountOptions.Get(1, "opt1")
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
