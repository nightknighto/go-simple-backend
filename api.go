package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func initializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	return router
}

func handleRoutes(r *mux.Router) {
	r.HandleFunc("/", homePage).Methods(http.MethodGet, http.MethodHead, http.MethodOptions)
	r.HandleFunc("/products", getAllProductsHandler).Methods(http.MethodGet, http.MethodHead, http.MethodOptions)
	r.HandleFunc("/products/{id}", getProductHandler).Methods(http.MethodGet, http.MethodHead, http.MethodOptions)
}

func startServer(port string, r *mux.Router) {
	log.Fatalln(http.ListenAndServe(port, r))
}

var counter int = 0

func homePage(w http.ResponseWriter, r *http.Request) {
	counter++
	log.Println("Endpoint hit: /")
	fmt.Fprintf(w, "Welcome visitor #%v", counter)
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

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: /products")
	products, e := getAllProducts(db)
	if e != nil {
		sendError(w, http.StatusInternalServerError, e.Error())
		return
	}

	sendResponse(w, http.StatusOK, products)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]

	parsedId, e := strconv.Atoi(idParam)
	if e != nil {
		sendError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Printf("The parsed ID is: %v", parsedId)

	prod, e := getProduct(db, parsedId)
	if e != nil {
		switch e {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, e.Error())
		}
		return
	}

	sendResponse(w, http.StatusOK, prod)
}
