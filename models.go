package main

import (
	"database/sql"
	"fmt"
	"log"
)

func initializeDatabase() *sql.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v", DbUser, DbPassword, DbName)

	db, e := sql.Open("mysql", connectionString)

	if e != nil {
		log.Fatalln(e)
	}

	return db
}

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

func getAllProducts(db *sql.DB) ([]Product, error) {
	results, e := db.Query("SELECT * FROM products")
	if e != nil {
		return nil, e
	}

	products := []Product{}

	for results.Next() {
		var row Product
		e = results.Scan(&row.Id, &row.Name, &row.Quantity, &row.Price)
		if e != nil {
			return nil, e
		}
		products = append(products, row)
	}

	e = results.Err()
	if e != nil {
		return nil, e
	}

	return products, nil
}

func getProduct(db *sql.DB, id int) (Product, error) {
	row := db.QueryRow("SELECT id, name, quantity, price FROM products WHERE id=?", id)
	var p Product
	e := row.Scan(&p.Id, &p.Name, &p.Quantity, &p.Price)
	if e != nil {
		return p, e
	}

	return p, nil
}

func createProduct(db *sql.DB, p *Product) (int, error) {
	result, e := db.Exec("INSERT INTO products(name, quantity, price) VALUES (?,?,?)", p.Name, p.Quantity, p.Price)
	if e != nil {
		return 0, e
	}

	insertedId, e := result.LastInsertId()
	if e != nil {
		return 0, e
	}

	return int(insertedId), nil
}
