package models

type BankAccount struct {
	ID           int     `json:"id" gorm:"primary_key"`
	Type string  `json:"type"`
	Name string `json:"name"`
	Description string `json:"description"`
	AccountNumber string `json:"account_number"`
	CustomerId *int `json:"customer_id"`
	Customer Customer `json:"customer"`
	ClientId int `json:"client_id"`
	TellerId int `json:"teller_id"`
	Client Client `json:"client"`
	Teller Teller `json:"teller"`
	InsertedAt string `json:"inserted_at"`
	UpdatedAt string `json:"updated_at"`
}
