package Model

type Account struct {
	Password string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Username string `json:"username"`
}