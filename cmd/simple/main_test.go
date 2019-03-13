package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestNoNulls(t *testing.T) {
	buf := &bytes.Buffer{}
	writeTable(buf, "economy", true, false)
	rows := bytes.Split(buf.Bytes(), []byte("\n"))
	for _, r := range rows {
		if bytes.Contains(r, []byte("XMAS")) {
			t.Errorf("XMAS present in list, even though contains nulls")
		}
	}
}

// this test connects to the db and reads known data into known columns.
// next test needs to load unknown data into unknown number of columns in a string format.
func TestConnectToMondial(t *testing.T) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, password, host, port, sslmode)
	fmt.Println(connStr)
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
		//log.Println(name, code, capital)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

//this learning test turns a database table into a csv
var done = false

func TestConnectToMondialGeneric(t *testing.T) {
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
	cols, err := rows.Columns()
	if err != nil {
		t.Fatalf("Columns: %v", err)
	}

	_vals := make([]sql.NullString, len(cols))
	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = &_vals[i]
	}

	for rows.Next() {
		err := rows.Scan(vals...)
		if err != nil {
			log.Fatal(err)
		}
		outs := make([]string, 0)
		for i := range cols {
			v := vals[i].(*sql.NullString)
			if v.Valid {
				outs = append(outs, strconv.Quote(v.String))
			} else {
				outs = append(outs, "")
			}
		}
		csv := strings.Join(outs, ",")
		if !done {
			fmt.Println(csv)
			done = true
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
