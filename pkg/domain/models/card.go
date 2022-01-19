package models

type Card struct {
	Id int `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Rarity int `json:"rarity"`

	ReleaseDate string `json:"releaseDate"`

	ImgRef string `json:"imgRef"`

	Collection string `json:"collection"`

	Qnt int `json:"qnt,omitempty"`

	Favorite bool `json:"favorite,omitempty"`
}
