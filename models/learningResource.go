package models

import (
	"fmt"
	"time"
)

// LearningResource is an element of the learning architecture
type LearningResource struct {
	ID              int64 `schema:"-"`
	Author          string
	Copyright       bool
	Licence         string
	Title           string
	Description     string
	Path            string    `sql:",unique"`
	AddedOn         time.Time `schema:"-"`
	LearningPoints  []*LearningPoint
	KnowledgePoints []*KnowledgePoint
	Comments        string
	Tags            []string
}

func (lr *LearningResource) String() string {
	text := ""

	tString := lr.AddedOn.Format("2006-01-02")

	text += fmt.Sprintf("\n%s, Author: %s (%s)",
		lr.Title,
		lr.Author,
		tString,
	)
	return text
}
