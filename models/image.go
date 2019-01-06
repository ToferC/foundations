package models

import (
	"time"
)

// Image  is the image and path for an Image
type Image struct {
	ID          int64 `schema:"-"`
	Title       string
	Description string
	Path        string    `schema:"-"`
	AddedOn     time.Time `schema:"-"`
}
