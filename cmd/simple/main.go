package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var user = os.Getenv("VIS_MONDIAL_USER")
var dbname = os.Getenv("VIS_MONDIAL_DBNAME")
var password = os.Getenv("VIS_MONDIAL_PASSWORD")
var host = os.Getenv("VIS_MONDIAL_HOST")
var port = os.Getenv("VIS_MONDIAL_PORT")
var sslmode = os.Getenv("VIS_MONDIAL_SSLMODE")

var flgServe = flag.Bool("serve", false, "serves the simle query service on specified port")
var flgPort = flag.String("port", "8080", "port to serve on, default is 8080")

func init() {}

func main() {
	flag.Parse()
	if *flgServe {
		http.HandleFunc("/mondial/names/", handleNames)
		http.HandleFunc("/mondial/", handle)
		log.Fatal(http.ListenAndServe("0.0.0.0:"+*flgPort, nil))
	}
	flag.Usage()
}

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	table := path.Base(r.URL.Path)
	h := false
	n := false
	if r.FormValue("h") == "true" {
		h = true
	}
	if r.FormValue("null") == "true" {
		n = true
	}
	writeTable(w, table, h, n)
}

func handleNames(w http.ResponseWriter, r *http.Request) {
	q := "SELECT table_name FROM information_schema.tables WHERE table_schema='public'"
	WriteQuery(w, q, false, false)
}

func writeTable(w io.Writer, table string, header bool, null bool) {
	WriteQuery(w, "SELECT * FROM "+table, header, null)
}

func WriteQuery(w io.Writer, q string, header bool, null bool) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, password, host, port, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if header {
		w.Write([]byte(strings.Join(cols, ",") + "\n"))
	}
	if err != nil {
		log.Fatalf("Columns: %v", err)
	}

	_vals := make([]sql.NullString, len(cols))
	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = &_vals[i]
	}

	hasOneInvalidCol := false
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
				if !null { //if a column is invalid exclude if null is set to false
					hasOneInvalidCol = true
				}
				outs = append(outs, "")
			}
		}
		if !hasOneInvalidCol {
			csv := strings.Join(outs, ",")
			w.Write([]byte(csv + "\n"))
		}
		//reset flag
		hasOneInvalidCol = false
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
