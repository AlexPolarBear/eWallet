package services

import (
	"encoding/json"
	"ewallet/internal/orm"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["walletId"]

	wallet, ok := orm.Wallets[ID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(wallet)
	fmt.Println(wallet)
	w.WriteHeader(http.StatusOK)
}
