package models

// KnowledgePoint is an element of the learning architecture
type KnowledgePoint struct {
	ID          int64 `schema:"-"`
	Name        string
	Description string
	Slug        string `schema:"-"`
	Tags        []*string
}
