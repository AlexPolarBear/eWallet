package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"ewallet/internal/dto"
)

func CreateWallet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()

		balance := 100.0

		wallet := dto.Wallet{
			ID:      id,
			Balance: balance,
		}

		_, err := db.Exec("INSERT INTO wallets (id, balance) VALUES ($1, $2)", wallet.ID, wallet.Balance)
		if err != nil {
			log.Fatal(err)
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
}

func getWalletFromDB(tx *sql.Tx, id string) (*dto.Wallet, error) {
	var wallet dto.Wallet

	err := tx.QueryRow("SELECT id, balance FROM wallets WHERE id = $1", id).
		Scan(&wallet.ID, &wallet.Balance)

	return &wallet, err

}

func SendMoney(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID := params["walletId"]

		ctx := context.Background()
		tx, err := db.BeginTx(ctx,
			&sql.TxOptions{Isolation: sql.LevelRepeatableRead, ReadOnly: false})
		if err != nil {
			log.Fatal(err)
		}
		defer tx.Rollback()

		fromWallet, err := getWalletFromDB(tx, ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Print(err)
			return
		}

		var requestData struct {
			To     string  `json:"to"`
			Amount float64 `json:"amount"`
		}

		err = json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		toWallet, err := getWalletFromDB(tx, requestData.To)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if fromWallet.Balance < requestData.Amount {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = tx.ExecContext(ctx, "UPDATE wallets SET balance = balance - $1 WHERE id = $2", requestData.Amount, fromWallet.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = tx.ExecContext(ctx, "UPDATE wallets SET balance = balance + $1 WHERE id = $2", requestData.Amount, toWallet.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = tx.ExecContext(ctx,
			"INSERT INTO transactions (time, wallet_from, wallet_to, amount) "+
				"VALUES ($1, $2, $3, $4)", time.Now(), fromWallet.ID, toWallet.ID, requestData.Amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tx.Commit()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func HistoryWallet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID := params["walletId"]

		transactions := []dto.Transactions{}

		rows, err := db.Query("SELECT time, wallet_from, wallet_to, amount "+
			"FROM transactions WHERE wallet_from = $1 OR wallet_to = $1", ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var tr dto.Transactions
			err := rows.Scan(&tr.Time, &tr.Wallet_from, &tr.Wallet_to, &tr.Amount)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			transactions = append(transactions, tr)
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(transactions)
		w.WriteHeader(http.StatusOK)
	}
}

func GetWallet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID := params["walletId"]

		var wallet dto.Wallet

		err := db.QueryRow("SELECT id, balance FROM wallets WHERE id = $1", ID).
			Scan(&wallet.ID, &wallet.Balance)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
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
}
