package main

import (
	"fmt"
	"log"
	"net/http"

	ser "ewallet/internal/services"
)

var port = 8000

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/wallet", ser.CreateWallet)
	mux.HandleFunc("/api/v1/wallet/{walletId}/send", ser.SendMoney)
	mux.HandleFunc("/api/v1/wallet/{walletId}/history", ser.HistoryWallet)
	mux.HandleFunc("/api/v1/wallet/{walletId}", ser.GetWallet)

	log.Printf("Connection on: http://localhost:%d \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
