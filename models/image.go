package models

import (
	"time"
)

// Image  is the image and path for an Image
type Image struct {
	ID          int
	Title       string
	Description string
	Path        string
	AddedOn     time.Time
}
