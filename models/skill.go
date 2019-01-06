package models

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
