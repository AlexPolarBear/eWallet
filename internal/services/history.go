package services

// func HistoryWallet(w http.ResponseWriter, r *http.Request) {
// params := mux.Vars(r)
// ID := params["walletId"]

// wallet, err := getWalletFromDB(ID)
// if err != nil {
// 	w.WriteHeader(http.StatusNotFound)
// 	return
// }

// transfers, err := getTransfersFromDB(ID)
// if err != nil {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	return
// }

// jsonTransfers := make([]map[string]interface{}, len(transfers))
// for i, transfer := range transfers {
// 	jsonTransfer := map[string]interface{}{
// 		"time":    transfer.Time.Format(time.RFC3339),
// 		"from":    transfer.From,
// 		"to":      transfer.To,
// 		"balance": transfer.Balance,
// 	}
// 	jsonTransfers[i] = jsonTransfer
// }

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	// json.NewEncoder(w).Encode(jsonTransfers)
// }
