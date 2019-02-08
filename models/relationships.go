package models

import "time"

// Relationship models connections between users on the system
type Relationship struct {
	SourceID, TargetID int64
	Type               string
	CreatedOn          time.Time
	DeletedOn          time.Time
}
