package models

import "fmt"

// Practice is an element of the learning architecture
type Practice struct {
	ID          int64 `schema:"-"`
	Name        string
	Description string
	Tags        []*string
	Skills      []*Skill
	Slug        string `schema:"-"`
}

func (p *Practice) String() string {

	text := fmt.Sprintf("%s %s",
		p.Name,
		p.Description)

	return text
}
