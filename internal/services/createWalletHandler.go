package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"ewallet/internal/orm"
)

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()

	balance := 100.0

	wallet := orm.Wallet{
		ID:      id,
		Balance: balance,
	}

	jsonBytes, err := json.Marshal(wallet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to marshal wallet:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonBytes)
	w.WriteHeader(http.StatusOK)
}
