package main

import (
	"database/sql"
	_ "encoding/json" // Import paket "encoding/json", namun tidak digunakan dalam kode ini
	"fmt"
	"log"
	"net/http"
	"time"

	handler "Assigment2/Handler"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Import driver PostgreSQL, namun tidak digunakan dalam kode ini
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres" // Sesuaikan dengan kata sandi PostgreSQL Anda
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

const PORT = ":8080"

func main() {
	// Memastikan koneksi ke database
	db, err = sql.Open("postgres", ConnectDbPsql(host, user, password, dbname, port))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database")

	// Mengatur handler CRUD
	r := mux.NewRouter()
	itemHandler := handler.NewItemHandler(db)

	// Mengatur endpoint untuk menangani permintaan terkait pesanan
	r.HandleFunc("/orders", itemHandler.ItemsHandler)
	r.HandleFunc("/orders/{id}", itemHandler.ItemsHandler)

	fmt.Println("Now listening on port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Fungsi untuk menghasilkan string koneksi database PostgreSQL
func ConnectDbPsql(host, user, password, name string, port int) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)
	return psqlInfo
}
