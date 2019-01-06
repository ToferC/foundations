package models

import (
	"fmt"
	"time"
)

//User implements a generic user model
type User struct {
	ID             int64  `schema:"-"`
	UserName       string `sql:",unique"`
	Email          string
	Password       string
	IsAdmin        bool            `schema:"-"`
	LearnerProfile *LearnerProfile `schema:"-"`
	Interests      *InterestMap    `schema:"-"`
	CreatedAt      time.Time       `sql:"default:now()" schema:"-"`
	UpdatedAt      time.Time       `schema:"-"`
}

func (u User) String() string {
	text := fmt.Sprintf("%s %s %s", u.UserName, u.Email, u.Password)
	return text
}
