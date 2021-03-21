package model

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
	Genres         []string           `json:"genres,omitempty"`
	//VoteCount        *int64            `json:"vote_count,omitempty"`
	OriginalLanguage *OriginalLanguage `json:"original_language,omitempty"`
	//OriginalTitle    *string           `json:"original_title,omitempty"`
	PosterPath       *string           `json:"poster_path,omitempty"`
	Title            *string           `json:"title,omitempty"`
	//Video            *bool             `json:"video,omitempty"`
	//VoteAverage      *float64          `json:"vote_average,omitempty"`
	ID               *int64            `json:"id,omitempty"`
	Overview         *string           `json:"overview,omitempty"`
	//Popularity       *float64          `json:"popularity,omitempty"`
	MediaType        *MediaType        `json:"media_type,omitempty"`
	//OriginalName     *string           `json:"original_name,omitempty"`
	OriginCountry    []string          `json:"origin_country,omitempty"`
	Name             *string           `json:"name,omitempty"`
	FirstAirDate     *string           `json:"first_air_date,omitempty"`
}

type MediaType string
const (
	Movie string = "movie"
	Tv string = "tv"
	All string = "all"
	Person string = "person"
)

type TimeWindow string
const (
	Day string = "day"
	Week = "week"
)

type OriginalLanguage string
const (
	En string = "en"
	Fr string = "fr"
	Ja string = "ja"
)

type Region string
const (
	CA string = "CA"
)