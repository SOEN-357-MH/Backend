package model

type Configuration struct {
	Images     *Images  `json:"images,omitempty"`
	ChangeKeys []string `json:"change_keys,omitempty"`
}

type Images struct {
	BaseURL       *string  `json:"base_url,omitempty"`
	SecureBaseURL *string  `json:"secure_base_url,omitempty"`
	BackdropSizes []string `json:"backdrop_sizes,omitempty"`
	LogoSizes     []string `json:"logo_sizes,omitempty"`
	PosterSizes   []string `json:"poster_sizes,omitempty"`
	ProfileSizes  []string `json:"profile_sizes,omitempty"`
	StillSizes    []string `json:"still_sizes,omitempty"`
}
