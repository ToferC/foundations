package models

// LearningPoint is an element of the learning architecture
type LearningPoint struct {
	ID          int64
	Name        string
	Description string
	Slug        string
	Tags        []*string
}
