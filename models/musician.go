package models

type Musician struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MusicianType string `json:"musician_type"`
}
