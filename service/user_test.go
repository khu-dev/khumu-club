package service

import (
	"github.com/khu-dev/khumu-club/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAdmin(t *testing.T) {
	user := &data.User{
		Groups: []string{
			"admin",
		},
	}
	assert.True(t, IsAdmin(user))

	userFromDto := data.MapUserDto2User(&data.UserDto{
		Username: "testuser",
		Groups: []*data.Group{
			&data.Group{
				ID:   1,
				Name: "admin",
			},
		},
	})

	assert.True(t, IsAdmin(userFromDto))
}
