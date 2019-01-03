package models

// Practice is an element of the learning architecture
type Practice struct {
	ID          int64
	Name        string
	Description string
	Skills      []*Skill
	Tags        []*string
	Slug        string
}
