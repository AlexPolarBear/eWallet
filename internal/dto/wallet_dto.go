package dto

type Wallet struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type Transactions struct {
	Time        string  `json:"time"`
	Wallet_from string  `json:"wallet_from"`
	Wallet_to   string  `json:"wallet_to"`
	Amount      float64 `json:"amount"`
}
