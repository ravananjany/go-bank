package models

type Account struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Balance string `json:"balance,omitempty"`
}

type TransferRequest struct {
	FromAccountID string `json:"from_account_id" binding:"required"`
	ToAccountID   string `json:"to_account_id" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
}

type DB struct {
	Store map[string]*Account `json:"store,omitempty"`
}
