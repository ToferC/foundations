package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/pg"
)

// Experience represents a learner event
type Experience struct {
	ID                 int64
	UserName           string
	Verb               string
	Stream             *Stream
	Practices          []*Practice
	Skills             []*Skill
	LearningResourceID int64
	LearningResource   *LearningResource
	OccurredAt         time.Time `schema:"-"`
	Validated          bool      `schema:"-"`
	Time               int
	Value              int
	Difficulty         int
	Points             int
	Depth              int
	Comments           string
	Tags               []string
}

func (e *Experience) String() string {

	tString := e.OccurredAt.Format("2006-01-02")
	text := ""

	text += fmt.Sprintf("\n[%s] %s | Topic: %s, Learning Resource: %s",
		tString,
		e.Verb,
		e.Stream.Name,
		e.LearningResource.Title,
	)

	if e.Practices != nil {
		for _, p := range e.Practices {
			text += fmt.Sprintf("%s | ", p.Name)
		}
	}

	return text
}

// GetExperience is an API for Experience
func (e *Experience) GetExperience(db *pg.DB) error {
	return errors.New("Not implemented")
}

// UpdateExperience is an API for Experience
func (e *Experience) UpdateExperience(db *pg.DB) error {
	return errors.New("Not Implemented")
}

// DeleteExperience is an API for Experience
func (e *Experience) DeleteExperience(db *pg.DB) error {
	return errors.New("Not Implemented")

}

// CreateExperience is an API for Experience
func (e *Experience) CreateExperience(db *pg.DB) error {
	return errors.New("Not Implemented")

}
