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
	account       *models.AccountModel
	accountOption *models.AccountOptionModel
	category      *models.CategoryModel
	unit          *models.UnitModel
	productGroup  *models.ProductGroupModel
	product       *models.ProductModel
}

type application struct {
	conf     *config.AppConfig
	log      *slog.Logger
	db       *database.DB
	dbModels dbModels
}

func main() {
	dbPool, err := database.NewFromConfig()
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
			&models.AccountOptionModel{DB: dbPool},
			&models.CategoryModel{DB: dbPool},
			&models.UnitModel{DB: dbPool},
			&models.ProductGroupModel{DB: dbPool},
			&models.ProductModel{DB: dbPool},
		},
	}

	product, err := app.dbModels.product.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("product: %#v", product))

	prodGroupId, err := app.dbModels.productGroup.Insert("Spezi")
	if err != nil {
		panic(err.Error())
	} else {
		app.log.Debug(fmt.Sprintf("Inserted product group with ID: %d", prodGroupId))
	}

	err = app.dbModels.unit.Insert("l")
	if err != nil {
		panic(err.Error())
	}

	account, err := app.dbModels.account.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("account: %#v", account))

	dummyAccount := models.CreateAccount("FirstName", "NickName", "LastName",
		"Email", "Phone", "12.34", 100, 1)
	accountId, err := app.dbModels.account.Insert(dummyAccount)
	if err != nil {
		panic(err.Error())
	} else {
		app.log.Debug(fmt.Sprintf("Inserted account with ID: %d", accountId))
	}

	categories, err := app.dbModels.category.Get(1)
	if err != nil {
		panic(err.Error())
	}
	app.log.Debug(fmt.Sprintf("category: %#v", categories))

	accountOptions, err := app.dbModels.accountOption.Get(1, "opt1")
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
