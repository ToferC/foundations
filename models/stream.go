package models

import "fmt"

// Stream is an element of the learning architecture
type Stream struct {
	ID   int64 `schema:"-"`
	Name string
	Tags []*string
	Slug string `schema:"-"`
}

func (s *Stream) String() string {

	text := fmt.Sprintf("%s",
		s.Name)

	return text
}
