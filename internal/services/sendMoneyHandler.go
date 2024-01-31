package services

import (
	"encoding/json"
	"ewallet/internal/orm"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SendMoney(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["walletId"]

	if _, ok := orm.Wallets[ID]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var requestData struct {
		To      string  `json:"to"`
		Balance float64 `json:"balance"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := orm.Wallets[requestData.To]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if orm.Wallets[ID].Balance < requestData.Balance {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// wallets[ID].Balance -= requestData.Balance
	// wallets[requestData.To].Balance += requestData.Balance

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Перевод выполнен успешно")
}
