package main

import (
	"encoding/json"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/models"
	"net/http"
	"strconv"
)

func (app *application) getAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	account, err := app.account.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		app.log.Error("Could not marshal account")

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		app.log.Error("Could not write response")
	}
}
