package Model

type account struct {
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"` 
	Email string `json:"email"`
	Username string `json:"username"`
}