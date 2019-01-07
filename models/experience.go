package models

import (
	"time"
)

// Experience represents a learner event
type Experience struct {
	ID             int64
	Subject        *User
	Verb           string
	Stream         *Stream
	Practice       *Practice
	Skill          *Skill
	LearningPoint  *[]LearningPoint
	KnowledgePoint *[]KnowledgePoint
	Noun           *LearningResource
	OccurredAt     time.Time `schema:"-"`
	Validated      bool      `schema:"-"`
	Time           int
	Value          int
	Difficulty     int
	Depth          int
}
