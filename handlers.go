package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var p Product

	e := json.NewDecoder(r.Body).Decode(&p)
	if e != nil {
		log.Println(e.Error())
		sendError(w, http.StatusBadRequest, "Malformed request body")
		return
	}

	id, e := createProduct(db, &p)
	if e != nil {
		sendError(w, http.StatusInternalServerError, e.Error())
		return
	}

	p.Id = id
	sendResponse(w, http.StatusOK, p)
}
