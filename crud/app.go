package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) init() error {

	app.Router = mux.NewRouter().StrictSlash(true)
	return nil
}

func (app *App) run() error {
	error := http.ListenAndServe("localhost:5000", app.Router)
	if error != nil {
		log.Fatal(error)
	}
	return nil
}

func SendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	produts, err := GetProducts()
	if err != nil {
		log.Panic("Some error ocuured")
	}
	SendResponse(w, http.StatusOK, produts)
}
func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products/", app.getProducts).Methods("GET")
}
