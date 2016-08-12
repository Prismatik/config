package users

import (
	"github.com/prismatik/jabba"
)

func Go() {
	users := []jabba.User{
		davidbanham,
		drewshowalter,
		larlyntanasap,
		nathanwinch,
		simontaylor,
	}
	for _, user := range users {
		jabba.AddUser(user)
	}
}
