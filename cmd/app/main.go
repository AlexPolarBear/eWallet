package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"ewallet/internal/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var port = 8000

func main() {
	connStr := "host=db port=5432 user=postgres password=postgres dbname=ewallet sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rout := mux.NewRouter()

	rout.HandleFunc("/api/v1/wallet", services.CreateWallet(db)).Methods("POST")
	rout.HandleFunc("/api/v1/wallet/{walletId}/send", services.SendMoney(db)).Methods("POST")
	rout.HandleFunc("/api/v1/wallet/{walletId}/history", services.HistoryWallet(db)).Methods("GET")
	rout.HandleFunc("/api/v1/wallet/{walletId}", services.GetWallet(db)).Methods("GET")

	log.Printf("Connection on: http://localhost:%d \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), rout))
}
