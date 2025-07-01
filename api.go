package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	return router
}

func handleRoutes(r *mux.Router) {
	r.HandleFunc("/products", getAllProductsHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/products/{id}", getProductHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/products", createProductHandler).Methods(http.MethodPost)
}

func startServer(port string, r *mux.Router) {
	log.Fatalln(http.ListenAndServe(port, r))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload any) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	err_msg := map[string]string{
		"error": err,
	}
	sendResponse(w, statusCode, err_msg)

}
