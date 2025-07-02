package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func initializeDatabase(DbUser string, DbPassword string, DbName string) *sql.DB {
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

func updateProduct(db *sql.DB, p Product) error {
	result, e := db.Exec("UPDATE products SET name=?, price=?, quantity=? WHERE id=?", p.Name, p.Price, p.Quantity, p.Id)
	if e != nil {
		return e
	}
	
	rowsAffected, e := result.RowsAffected()
	if e != nil {
		return e
	}
	
	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func deleteProduct(db *sql.DB, id int) error {
	result, e := db.Exec("DELETE FROM products WHERE id=?", id)
	if e != nil {
		return e
	}

	rowsAffected, e := result.RowsAffected()
	if e != nil {
		return e
	}
	
	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}