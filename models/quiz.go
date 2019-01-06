package models

import (
	"time"
)

type Quiz struct {
	ID               int64
	Title            string
	Questions        []*Question
	Tagline          string
	Body             string
	SubTitle         string
	Author           *User
	Images           []*Image
	Videos           []*Video
	Likes            int        `schema:"-"`
	PublishedOn      *time.Time `schema:"-"`
	EditedOn         []*Edit    `schema:"-"`
	Tags             []*string
	LearningPoints   []*LearningPoint
	DigitalStandards []string
	BannerImage      *Image `schema:"-"`
	slug             string `schema:"-"`
}

type Question struct {
	Question string
	Answer   int
}
