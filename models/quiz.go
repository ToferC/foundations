package models

import (
	"time"
)

// Quiz represents a test of a user's knowledge
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

// Question is an question in a quiz
type Question struct {
	Question      string
	Answers       []string
	CorrectAnswer int
}
