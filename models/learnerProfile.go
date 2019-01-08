package models

import (
	"fmt"
)

// LearnerProfile represents a learner's experiences
type LearnerProfile struct {
	ID int64

	Experiences []*Experience

	slug string
}

func (lp *LearnerProfile) String() {
	fmt.Println(lp)
}
