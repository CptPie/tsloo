package models

import (
	"errors"
	"fmt"
	"time"
)

type Vote struct {
	Id      int
	User    *User
	Song    *Song
	Rating  int
	Comment string
	Added   time.Time
	Updated time.Time
}

func (v Vote) Update(u *User, rating int) (bool, error) {
	if u != v.User {
		return false, errors.New(fmt.Sprintf("User %s is not authorized to change vote %d", u.Name, v.Id))
	}

	v.Rating = rating
	return true, nil
}
