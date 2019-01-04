package models

import (
	"time"
)

// Video  is the image and path for an Image
type Video struct {
	ID             int
	Title          string
	Description    string
	Path           string
	LearningPoints []*LearningPoint
	AddedOn        time.Time
}
