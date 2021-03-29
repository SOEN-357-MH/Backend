package Model

type Account struct {
	Password       string `json:"password,omitempty"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	MovieWatchlist []int  `json:"movieWatchlist,omitempty" bson:"movie_watchlist"`
	ShowWatchlist  []int  `json:"showWatchlist,omitempty" bson:"show_watchlist"`
}

type Result struct {
	Page         *int64  `json:"page,omitempty"`
	Results      []Media `json:"results,omitempty"`
	TotalPages   *int64  `json:"total_pages,omitempty"`
	TotalResults *int64  `json:"total_results,omitempty"`
}

type Media struct {
	ReleaseDate      *string           `json:"release_date,omitempty"`
	Adult            *bool             `json:"adult,omitempty"`
	BackdropPath     *string           `json:"backdrop_path,omitempty"`
	GenreIDS         []int64           `json:"genre_ids,omitempty"`
	Genres           []string          `json:"genres,omitempty"`
	OriginalLanguage *OriginalLanguage `json:"original_language,omitempty"`
	PosterPath       *string           `json:"poster_path,omitempty"`
	Title            *string           `json:"title,omitempty"`
	ID               *int64            `json:"id,omitempty"`
	Overview         *string           `json:"overview,omitempty"`
	OriginCountry    []string          `json:"origin_country,omitempty"`
	Name             *string           `json:"name,omitempty"`
	FirstAirDate     *string           `json:"first_air_date,omitempty"`
}

type OriginalLanguage string

const (
	En string = "en"
	Fr string = "fr"
	Ja string = "ja"
)
