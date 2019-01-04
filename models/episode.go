package models

import (
	"time"
)

// Episode represents a Busrides Episode
type Episode struct {
	ID               int64
	Title            string
	Tagline          string
	Body             string
	SubTitle         string
	Author           *User
	Image            *Image
	Videos           []*Video
	Likes            int
	PublishedOn      *time.Time
	EditedOn         []*Edit
	Tags             []*string
	LearningPoints   []*LearningPoint
	DigitalStandards []string
	BannerImage      *Image
	slug             string
}

// Edit represents an edit to content
type Edit struct {
	Description string
	Date        *time.Time
}
