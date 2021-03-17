package Model

import "github.com/go-bongo/bongo"

type Account struct {
	bongo.DocumentBase `bson:",inline"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Username string `json:"username"`
}