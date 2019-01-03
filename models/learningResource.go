package models

import (
	"time"
)

// LearningResource is an element of the learning architecture
type LearningResource struct {
	ID          int64
	Author      string
	Copyright   bool
	Licence     string
	Title       string
	Description string
	Path        string
	AddedOn     time.Time
}
