package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var user = os.Getenv("VIS_MONDIAL_USER")
var dbname = os.Getenv("VIS_MONDIAL_DBNAME")
var password = os.Getenv("VIS_MONDIAL_PASSWORD")
var host = os.Getenv("VIS_MONDIAL_HOST")
var port = os.Getenv("VIS_MONDIAL_PORT")
var sslmode = os.Getenv("VIS_MONDIAL_SSLMODE")

// this test connects to the db and reads known data into known columns.
// next test needs to load unknown data into unknown number of columns in a string format.
func TestConnectToMondial(t *testing.T) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, password, host, port, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM country")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var name, code, area, destination string
	var capital, province sql.NullString
	for rows.Next() {
		err := rows.Scan(&name, &code, &capital, &province, &area, &destination)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, code, capital)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}