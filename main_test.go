package main

import (
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

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 29.99, 50)
	request := httptest.NewRequest("GET", "/products/1", nil)
	recorder := httptest.NewRecorder()

	testRouter.ServeHTTP(recorder, request)

	log.Println("Returned Code: ", recorder.Code)
	log.Println("Returned Body: ", recorder.Body)
	expectedStatusCode := http.StatusOK

	if expectedStatusCode != recorder.Code {
		t.Errorf("Expected status: %v, but received: %v\n", expectedStatusCode, recorder.Code)
	}
}
