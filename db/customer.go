package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	connStr := "host=localhost port=8080 dbname=shope user=postgres password=Pawan@2003 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func CreateCustomerTable(Db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS customer(
		id SERIAL PRIMARY KEY,
		name VARCHAR (100) NOT NULL,
		mobile_no VARCHAR(100) NOT NULL,
		address Text,
		created TIMESTAMP DEFAULT NOW ()
	)`
	_, err := Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// func insertProduct(db *sql.DB, product Product) int {
// 	query := `INSERT INTO product (name,price,available)
// 	VALUES($1,$2,$3) RETURNING id`

// 	var pk int
// 	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return pk

// }
