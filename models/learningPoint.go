package models

// LearningPoint is an element of the learning architecture
type LearningPoint struct {
	ID          int64 `schema:"-"`
	Name        string
	Description string
	Slug        string `schema:"-"`
	Tags        []*string
	Difficulty  int
}
