package service

import (
	"github.com/AlekSi/pointer"
	"github.com/khu-dev/khumu-club/data"
	"github.com/khu-dev/khumu-club/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewTestClubService() *ClubService {
	return &ClubService{DB: repository.ConnectForTest()}
}

func TestClubService_CreateClub(t *testing.T) {
	service := NewTestClubService()
	defer service.DB.Close()

	hashTags := []string{"개발", "스타트업", "디자인"}
	body := &data.ClubDto{
		Name:        pointer.ToString("쿠뮤 개발 동아리"),
		Summary:     pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Description: pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다.이런 저런 서비스를 개발해나가는 동아리 입니다.\n\n이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Email:       pointer.ToString("dev.umijs@gmail.com"),
		Hashtags:    hashTags,
	}
	newClub, err := service.CreateClub(body)
	assert.NoError(t, err)
	assert.Len(t, newClub.Hashtags, 3)
	for i, tag := range newClub.Hashtags {
		assert.Equal(t, hashTags[i], tag)
	}
}

func TestClubService_ListClub(t *testing.T) {
	service := NewTestClubService()
	defer service.DB.Close()

	hashTags := []string{"개발", "스타트업", "디자인"}
	body := &data.ClubDto{
		Name:        pointer.ToString("쿠뮤 개발 동아리"),
		Summary:     pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Description: pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다.이런 저런 서비스를 개발해나가는 동아리 입니다.\n\n이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Email:       pointer.ToString("dev.umijs@gmail.com"),
		Hashtags:    hashTags,
	}
	_, err := service.CreateClub(body)
	assert.NoError(t, err)

	results, err := service.ListClub()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}
