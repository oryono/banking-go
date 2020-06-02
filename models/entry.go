package models

type Entry struct {
	ID           int     `json:"id" gorm:"primary_key"`
	Type string  `json:"type"`
	TransactionReference  string `json:"transaction_reference"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
	AccountId int `json:"account_id"`
	Account BankAccount `json:"account"`
	RunningBalance float64 `json:"running_balance"`
	InsertedAt string `json:"inserted_at"`
	UpdatedAt string `json:"updated_at"`
}
