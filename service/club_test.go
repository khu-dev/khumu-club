package service

import (
	"context"
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

	categories := []string{"개발", "스타트업", "디자인"}
	for _, ctg := range categories {
		_, err := service.DB.Category.Create().
			SetID(ctg).
			Save(context.TODO())
		assert.NoError(t, err)
	}

	body := &data.ClubDto{
		Name:        pointer.ToString("쿠뮤 개발 동아리"),
		Summary:     pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Description: pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다.이런 저런 서비스를 개발해나가는 동아리 입니다.\n\n이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Email:       pointer.ToString("dev.umijs@gmail.com"),
		Categories:  categories,
	}
	newClub, err := service.CreateClub(body)
	assert.NoError(t, err)
	assert.Len(t, newClub.Categories, 3)
	for _, category := range newClub.Categories {
		assert.Contains(t, categories, category)
	}
}

func TestClubService_ListClub(t *testing.T) {
	service := NewTestClubService()
	defer service.DB.Close()

	categories := []string{"개발", "스타트업", "디자인"}
	for _, ctg := range categories {
		_, err := service.DB.Category.Create().
			SetID(ctg).
			Save(context.TODO())
		assert.NoError(t, err)
	}
	body := &data.ClubDto{
		Name:        pointer.ToString("쿠뮤 개발 동아리"),
		Summary:     pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Description: pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다.이런 저런 서비스를 개발해나가는 동아리 입니다.\n\n이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Email:       pointer.ToString("dev.umijs@gmail.com"),
		Categories:  categories,
	}
	_, err := service.CreateClub(body)
	assert.NoError(t, err)

	results, err := service.ListClubs()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestClubService_ListClub_소트(t *testing.T) {
	service := NewTestClubService()
	defer service.DB.Close()

	categories := []string{"개발", "스타트업", "디자인"}
	for _, ctg := range categories {
		_, err := service.DB.Category.Create().
			SetID(ctg).
			Save(context.TODO())
		assert.NoError(t, err)
	}
	defaultBody := data.ClubDto{
		Name:        pointer.ToString("쿠뮤 개발 동아리"),
		Summary:     pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Description: pointer.ToString("이런 저런 서비스를 개발해나가는 동아리 입니다.이런 저런 서비스를 개발해나가는 동아리 입니다.\n\n이런 저런 서비스를 개발해나가는 동아리 입니다."),
		Email:       pointer.ToString("dev.umijs@gmail.com"),
		Categories:  categories,
	}
	for i := 0; i < 10; i++ {
		body := defaultBody
		if i < 3 {
			body.Recommended = pointer.ToBool(true)
		}
		_, err := service.CreateClub(&body)
		assert.NoError(t, err)
	}

	results, err := service.ListClubs()
	assert.NoError(t, err)
	for i := 0; i < 3; i++ {
		assert.True(t, *results[i].Recommended)
	}

	results2, err := service.ListClubs()
	assert.NoError(t, err)
	isAllSame := true

	for i := 0; i < 10; i++ {
		if results[i].ID != results2[i].ID {
			isAllSame = false
			break
		}
	}
	assert.False(t, isAllSame, "동아리 List 시마다 랜덤해야합니다. 아주 희박한 확률로 우연의 일치일 수 있습니다.")
}
