package Model

type Account struct {
	Password string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Username string `json:"username"`
	MovieWatchlist []int `json:"movieWatchlist,omitempty" bson:"movie_watchlist"`
	ShowWatchlist []int `json:"showWatchlist,omitempty" bson:"show_watchlist"`

}