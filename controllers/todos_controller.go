package controllers

import (
	"todo_api_go/models"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	countories := make([]models.Country, 2)
	country1 := models.Country{Code: 1, Name: "Japan"}
	country2 := models.Country{Code: 2, Name: "USA"}
	countories[0] = country1
	countories[1] = country2

	w.WriteJson(&countories)
}
