package models

type Client struct {
	ID           int     `json:"id" gorm:"primary_key"`
	Name string  `json:"name"`
	Phone  string `json:"phone"`
	Address string `json:"address"`
	BankAccount *BankAccount `json:"bank_account" gorm:"ForeignKey:account_id"`
	InsertedAt string `json:"inserted_at"`
	UpdatedAt string `json:"updated_at"`
}
