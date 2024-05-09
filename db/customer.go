package db

import (
	"database/sql"
	"fmt"
	"log"
	"shope/modle"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connStr := "host=localhost port=8080 dbname=shope user=postgres password=Pawan@2003 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to Database successfully..")
	return db
}

func CreateCustomerTable(Db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS customers(
        id SERIAL PRIMARY KEY,
        name VARCHAR (100) NOT NULL,
        mobile_no TEXT NOT NULL,
        address TEXT,
        product TEXT
        )`
	_, err := Db.Exec(query)
	if err != nil {
		log.Fatalf("error in executing query %v", err)
	}

}
func InsertCustomer(customer modle.Customer) int64 {

	db := ConnectDb()
	defer db.Close()
	CreateCustomerTable(db)
	sqlStatement := `INSERT INTO customers (name, mobile_no,address, product) VALUES ($1, $2, $3,$4) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, customer.Name, customer.Mobile, customer.Address, customer.Product).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}
func GetCustomerById(id int64) (modle.Customer, error) {
	db := ConnectDb()
	defer db.Close()
	var customer modle.Customer
	sqlStmt := `SELECT * FROM customers WHERE id=$1`
	row := db.QueryRow(sqlStmt, id)
	err := row.Scan(&customer.Id, &customer.Name, &customer.Mobile, &customer.Address, &customer.Product)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now Rows were returend!")
		return customer, nil
	case nil:
		return customer, nil
	default:
		log.Fatalf("Unable to scan row %v", err)
	}
	return customer, err
}
func GetAllCustomers() ([]modle.Customer, error) {
	db := ConnectDb()
	defer db.Close()
	var customers []modle.Customer
	sqlStmt := `SELECT * FROM customers `
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatalf("Unable TO execute the query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var customer modle.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Mobile, &customer.Address, &customer.Product)
		if err != nil {
			log.Fatalf("Unable to scan the ro %v", err)
		}
		customers = append(customers, customer)
	}
	return customers, err
}

func UpdateById(id int64, customer modle.Customer) int64 {
	db := ConnectDb()

	defer db.Close()

	sqlStatement := `UPDATE customers SET name=$2, mobile_no=$3, address=$4,product=$5 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, customer.Name, customer.Mobile, customer.Address, customer.Product)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
func DeleteCustomerById(id int64) int64 {
	db := ConnectDb()
	defer db.Close()
	sqlStmt := `DELETE FROM customers WHERE id=$1 `
	res, err := db.Exec(sqlStmt, id)
	if err != nil {
		log.Fatalf("Unable TO execute the query %v", err)
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows.%v", err)
	}
	fmt.Printf("Total row affected.%v", rowAffected)
	return rowAffected
}
