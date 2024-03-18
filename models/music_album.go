package models

import "time"

type MusicAlbum struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ReleaseDate time.Time  `json:"release_date"`
	Genre       string     `json:"genre"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	Musicians   []Musician `json:"musicians"`
}
