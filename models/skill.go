package models

import "fmt"

// Skill is an element of the learning architecture
type Skill struct {
	ID              int64 `schema:"-"`
	Name            string
	Description     string
	LearningPoints  []*LearningPoint
	KnowledgePoints []*KnowledgePoint
	Slug            string `schema:"-"`
	Tags            []*string
}

func (s *Skill) String() string {

	text := fmt.Sprintf("%s %s",
		s.Name,
		s.Description)

	return text
}
