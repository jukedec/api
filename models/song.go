package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Song is used by pop to map your songs database table to your go code.
type Song struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	ImgUrl      string    `json:"img_url" db:"img_url"`
	SongUrl     string    `json:"song_url" db:"song_url"`
	Price       float32   `json:"price" db:"price"`
	ArtistID    uuid.UUID `json:"artist_id" db:"artist_id"`
	AlbumID     uuid.UUID `json:"album_id" db:"album_id"`
}

// String is not required by pop and may be deleted
func (s Song) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Songs is not required by pop and may be deleted
type Songs []Song

// String is not required by pop and may be deleted
func (s Songs) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Song) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Song) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Song) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
