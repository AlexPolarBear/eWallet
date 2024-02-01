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
	connStr := "host=localhost port=5432 user=postgres " +
		"password=password dbname=ewallet"

	db, err := sql.Open("postgres", connStr)
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

/*
    //create router
    router := mux.NewRouter()
    router.HandleFunc("/users", getUsers(db)).Methods("GET")
    router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
    router.HandleFunc("/users", createUser(db)).Methods("POST")
    router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
    router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

    //start server
    log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}*/
