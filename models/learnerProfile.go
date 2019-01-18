package models

import (
	"fmt"
)

// LearnerProfile represents a learner's experiences
type LearnerProfile struct {
	ID int64

	Experiences []*Experience

	// LearningTargets and Current Learnings are maps of a year 2019 to a target 10000pts
	LearningTargets map[string][]int
	CurrentYear     string

	slug string
}

func (lp *LearnerProfile) String() {
	fmt.Println()
}
