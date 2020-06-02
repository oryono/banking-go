package models

type Teller struct {
	ID           int     `json:"id" gorm:"primary_key"`
	Name string  `json:"name"`
	Email  string `json:"phone"`
	ClientId int `json:"client_id"`
	Client Client `json:"client" gorm:"ForeignKey:client_id"`
	InsertedAt string `json:"inserted_at"`
	UpdatedAt string `json:"updated_at"`
}
