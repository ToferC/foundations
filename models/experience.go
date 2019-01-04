package models

import (
	"time"
)

// Experience represents a learner event
type Experience struct {
	Subject        *User
	Verb           string
	Skill          *Skill
	Practice       *Practice
	LearningPoint  *LearningPoint
	KnowledgePoint *KnowledgePoint
	Noun           *LearningResource
	OccurredAt     time.Time
	Validated      bool
	Difficulty     int
	Depth          int
}
