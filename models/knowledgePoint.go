package models

// KnowledgePoint is an element of the learning architecture
type KnowledgePoint struct {
	ID          int64
	Name        string
	Description string
	Slug        string
	Tags        []*string
}
