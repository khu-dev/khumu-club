package service

import "github.com/khu-dev/khumu-club/data"

func IsAdmin(user *data.User) bool {
	for _, group := range user.Groups {
		if group == "admin" {
			return true
		}
	}

	return false
}
