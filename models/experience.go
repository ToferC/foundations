package models

import (
	"time"
)

// Experience represents a learner event
type Experience struct {
	ID             int64
	Subject        *User
	Verb           string
	Skill          *Skill
	Practice       *Practice
	LearningPoint  *LearningPoint
	KnowledgePoint *KnowledgePoint
	Noun           *LearningResource
	OccurredAt     time.Time `schema:"-"`
	Validated      bool      `schema:"-"`
	Difficulty     int
	Depth          int
}
