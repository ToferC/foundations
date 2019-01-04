package models

import (
	"time"
)

type Quiz struct {
	Title            string
	Questions        []*Question
	Tagline          string
	Body             string
	SubTitle         string
	Author           *User
	Images           []*Image
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

type Question struct {
	Question string
	Answer   int
}
