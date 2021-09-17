package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/khu-dev/khumu-club/adapter/slack"
	"github.com/khu-dev/khumu-club/data"
	"github.com/khu-dev/khumu-club/ent"
	"github.com/khu-dev/khumu-club/ent/category"
	"github.com/khu-dev/khumu-club/ent/club"
	"github.com/khu-dev/khumu-club/util"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

var (
	ErrNoRequiredField = errors.New("필수 입력 필드를 입력하지 않았습니다")
)

type ClubService struct {
	DB           *ent.Client
	SlackAdapter slack.SlackAdapter
}

func (s *ClubService) CreateClub(body *data.ClubDto) (*data.ClubDto, error) {
	log.Info("새로운 동아리를 생성합니다.")
	defer log.Info("새로운 동아리 생성을 마칩니다.")
	if body.Name == nil ||
		body.Summary == nil ||
		body.Description == nil {
		log.Error(ErrNoRequiredField)
		return nil, ErrNoRequiredField
	}
	if body.Categories == nil {
		body.Categories = []string{}
	}
	if body.Images == nil {
		body.Images = []string{}
	}
	club, err := s.DB.Club.Create().
		SetName(*body.Name).
		SetSummary(*body.Summary).
		SetDescription(*body.Description).
		AddCategoryIDs(body.Categories...).
		SetImages(body.Images).
		SetNillableHomepage(body.Homepage).
		SetNillableInstagram(body.Instagram).
		SetNillableInstagram(body.Instagram).
		SetNillableFacebook(body.Facebook).
		SetNillablePhone(body.Phone).
		SetNillableEmail(body.Email).
		SetNillableRecommended(body.Recommended).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}
	club.Edges.Categories = club.QueryCategories().AllX(context.TODO())
	return data.MapClubToClubDto(club), err
}
func (s *ClubService) ListClubs() ([]*data.ClubDto, error) {
	results := s.DB.Club.Query().WithCategories().AllX(context.TODO())
	sortedResults := s.SortClubList(results)
	outputs := make([]*data.ClubDto, len(sortedResults))
	for i, club := range sortedResults {
		outputs[i] = data.MapClubToClubDto(club)
	}

	return outputs, nil
}

func (s *ClubService) ListClubsByCategoryContaining(ctg string) ([]*data.ClubDto, error) {
	log.Info("카테고리가 포함된 동아리 검색")
	results, err := s.DB.Club.Query().Where(club.HasCategoriesWith(
		category.ID(ctg)),
	).
		WithCategories().
		All(context.TODO())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	sortedResults := s.SortClubList(results)
	outputs := make([]*data.ClubDto, len(sortedResults))
	for i, club := range sortedResults {
		outputs[i] = data.MapClubToClubDto(club)
	}

	return outputs, nil
}

// 단조로운 동아리 리스트를 없애기 위해 순서를 섞어준다.
func (s *ClubService) SortClubList(src []*ent.Club) []*ent.Club {
	recommendedClubs := make([]*ent.Club, 0)
	normalClubs := make([]*ent.Club, 0)
	for _, club := range src {
		if club.Recommended {
			recommendedClubs = append(recommendedClubs, club)
		} else {
			normalClubs = append(normalClubs, club)
		}
	}

	output := make([]*ent.Club, len(src))
	for i, rec := range recommendedClubs {
		output[i] = rec
	}
	for i, randomIndex := range rand.Perm(len(normalClubs)) {
		output[len(recommendedClubs)+i] = normalClubs[randomIndex]
	}

	return output
}

func (s *ClubService) ClubAddRequest(ctx *fiber.Ctx, body *data.ClubAddOrModifyRequestDto) error {
	user := util.GetRequestUser(ctx)
	err := s.SlackAdapter.SendMessage("[동아리 추가 요청]"+body.Title, user.ID, body.Content)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *ClubService) ClubModifyRequest(ctx *fiber.Ctx, body *data.ClubAddOrModifyRequestDto) error {
	user := util.GetRequestUser(ctx)
	err := s.SlackAdapter.SendMessage("[동아리 수정 요청]"+body.Title, user.ID, body.Content)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *ClubService) ListCategories(ctx *fiber.Ctx) []string {
	// 그냥 코드 자체로 간단히 관리함.
	ctgs, err := s.DB.Category.Query().All(context.TODO())
	if err != nil {
		log.Error(err)
		return []string{}
	}

	ctgStrings := make([]string, 0)
	for _, ctg := range ctgs {
		ctgStrings = append(ctgStrings, ctg.ID)
	}

	return ctgStrings
}
