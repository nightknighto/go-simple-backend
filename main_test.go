package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// var db *sql.DB
var testRouter *mux.Router

func TestMain(m *testing.M) {
	db = initializeDatabase("root", "root", "test")
	testRouter = initializeRouter()

	createTable()

	m.Run()
}

func createTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		quantity INT,
		price FLOAT
		);`)

	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	_, err := db.Exec("TRUNCATE TABLE products")

	if err != nil {
		log.Fatal(err)
	}
}

func addProduct(name string, price float32, quantity int) {
	_, e := db.Exec("INSERT INTO products(name, price, quantity) VALUES (?,?,?)", name, price, quantity)
	if e != nil {
		log.Println(e)
	}
}

func checkStatusCode(t *testing.T, expectedStatusCode int, receivedStatusCode int) {
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status: %v, but received: %v\n", expectedStatusCode, receivedStatusCode)
	}
}

func sendRequest(method string, target string, body io.Reader) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, target, body)
	recorder := httptest.NewRecorder()

	testRouter.ServeHTTP(recorder, request)

	return recorder
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 29.99, 50)

	response := sendRequest("GET", "/products/1", nil)

	log.Println("Returned Code: ", response.Code)
	log.Println("Returned Body: ", response.Body)

	checkStatusCode(t, http.StatusOK, response.Code)
}

func TestCreateProduct(t *testing.T) {
	clearTable()
	var product = `{"name":"chair","quantity":4,"price":29.99}`

	response := sendRequest("POST", "/products", bytes.NewBufferString(product))

	log.Println("Returned Code: ", response.Code)
	log.Println("Returned Body: ", response.Body)

	var parsedBody map[string]any
	json.Unmarshal(response.Body.Bytes(), &parsedBody)

	checkStatusCode(t, http.StatusCreated, response.Code)

	if parsedBody["name"] != "chair" {
		t.Errorf("Expected name: %v, Got: %v", "chair", parsedBody["name"])
	}

	if parsedBody["quantity"] != 4.0 {
		t.Errorf("Expected quantity: %v, Got: %v", 4, parsedBody["quantity"])
	}

	if parsedBody["price"] != 29.99 {
		t.Errorf("Expected price: %v, Got: %v", 29.99, parsedBody["price"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()

	pName, pQuantity, pPrice := "chair", 5, 39.99
	addProduct(pName, float32(pPrice), pQuantity)

	response := sendRequest("GET", "/products/1", nil)

	checkStatusCode(t, http.StatusOK, response.Code)

	response = sendRequest("DELETE", "/products/1", nil)

	checkStatusCode(t, http.StatusOK, response.Code)

	response = sendRequest("GET", "/products/1", nil)

	checkStatusCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	clearTable()

	pName, pQuantity, pPrice := "chair", 5, 39.99
	addProduct(pName, float32(pPrice), pQuantity)

	// 1- GET the old price

	response := sendRequest("GET", "/products/1", nil)

	var b map[string]any
	json.Unmarshal(response.Body.Bytes(), &b)

	checkStatusCode(t, http.StatusOK, response.Code)

	if b["price"] != pPrice {
		t.Errorf("Expected price: %v, Got: %v", pPrice, b["price"])
	}

	// 2- Change the Price
	newPrice := 79.99

	requestBody := fmt.Sprintf(`{"name":"%v", "quantity":%v, "price": %v}`, pName, pQuantity, newPrice)

	response = sendRequest("PUT", "/products/1", bytes.NewBufferString(requestBody))

	b = map[string]any{}
	json.Unmarshal(response.Body.Bytes(), &b)

	checkStatusCode(t, http.StatusOK, response.Code)

	if b["price"] != newPrice {
		t.Errorf("Expected price: %v, Got: %v", newPrice, b["price"])
	}

	// 3- GET the new price

	response = sendRequest("GET", "/products/1", nil)

	b = map[string]any{}
	json.Unmarshal(response.Body.Bytes(), &b)

	checkStatusCode(t, http.StatusOK, response.Code)

	if b["price"] != newPrice {
		t.Errorf("Expected price: %v, Got: %v", newPrice, b["price"])
	}
}
