package models

// Stream is an element of the learning architecture
type Stream struct {
	ID        int64 `schema:"-"`
	Name      string
	Practices []*Practice
	Tags      []*string
	Slug      string `schema:"-"`
}
