package models

// InterestMap is a record of user-created interests
type InterestMap struct {
	Streams   []*Stream
	Practices []*Practice
}
