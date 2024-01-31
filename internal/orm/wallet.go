package orm

type Wallet struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

var Wallets map[string]Wallet
