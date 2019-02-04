package models

import (
	"time"
)

// Episode represents a Busrides Episode
type Episode struct {
	ID                int64 `schema:"-"`
	Title             string
	Tagline           string
	SubTitle          string
	Body              string
	Author            *User `schema:"-"`
	Image             *Image
	Videos            []*Video
	Podcasts          []*Podcast
	Likes             int       `schema:"-"`
	PublishedOn       time.Time `schema:"-"`
	Edits             []*Edit   `schema:"-"`
	Tags              []string
	LearningPoints    []*LearningPoint
	LearningResources []*LearningResource
	Experience        *Experience
	DigitalStandards  []string
	BannerImage       *Image
	Slug              string `schema:"-"`
}

// Edit represents an edit to content
type Edit struct {
	Description string
	Date        *time.Time
}
