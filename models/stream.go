package models

import "fmt"

// Stream is an element of the learning architecture
type Stream struct {
	ID              int64 `schema:"-"`
	Name            string
	Description     string
	Image           *Image
	Tags            []*string
	Practices       map[string]*Practice
	LearningTargets map[string][]int
	Selected        bool
	Expertise       int
	Slug            string `schema:"-"`
}

func (s *Stream) String() string {

	text := fmt.Sprintf("Stream: %s - Practices: ",
		s.Name)

	for _, v := range s.Practices {
		text += fmt.Sprintf("%s | ", v.Name)
	}

	return text
}
