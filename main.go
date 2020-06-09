package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"todo_api_go/controllers"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "app"
	password = "app-pass"
	dbname   = "todo_development"
)

var db *sql.DB

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/coutries", controllers.GetAllCountries),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// db
	dbsetup()

	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func dbsetup() {
	psqlInfo := fmt.Sprintf(
		"host%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connected!")
	}
	defer db.Close()
}
