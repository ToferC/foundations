package models

// Stream is an element of the learning architecture
type Stream struct {
	ID        int64
	Name      string
	Practices []*Practice
	Tags      []*string
	Slug      string
}
