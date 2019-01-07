package models

import (
	"time"
)

// Episode represents a Busrides Episode
type Episode struct {
	ID               int64 `schema:"-"`
	Title            string
	Tagline          string
	Body             string
	SubTitle         string
	Author           *User `schema:"-"`
	Image            *Image
	Videos           []*Video
	Podcasts         []*Podcast
	Likes            int       `schema:"-"`
	PublishedOn      time.Time `schema:"-"`
	Edits            []*Edit   `schema:"-"`
	Tags             []string
	LearningPoints   []*LearningPoint
	DigitalStandards []string
	BannerImage      *Image
	slug             string `schema:"-"`
}

// Edit represents an edit to content
type Edit struct {
	Description string
	Date        *time.Time
}
