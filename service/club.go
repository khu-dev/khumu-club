package service

import (
    "context"
    "errors"
    "github.com/khu-dev/khumu-club/data"
    "github.com/khu-dev/khumu-club/ent"
    log "github.com/sirupsen/logrus"
)

var (
	ErrNoRequiredField = errors.New("필수 입력 필드를 입력하지 않았습니다")
)

type ClubService struct{
    DB *ent.Client
}

func (s *ClubService) CreateClub(body *data.ClubDto) (*data.ClubDto, error) {
    if body.Name == nil ||
        body.Summary == nil ||
        body.Description == nil{
        log.Error(ErrNoRequiredField)
        return nil, ErrNoRequiredField
    }
    if body.Hashtags == nil {
        body.Hashtags = []string{}
    }
    if body.Images == nil {
        body.Images = []string{}
    }
    club, err := s.DB.Club.Create().
        SetName(*body.Name).
        SetSummary(*body.Summary).
        SetDescription(*body.Description).
        SetHashtags(body.Hashtags).
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

    return data.MapClubToClubDto(club), err
}
func (s *ClubService) ListClub() ([]*data.ClubDto, error) {
    results := s.DB.Club.Query().AllX(context.TODO())
    outputs := make([]*data.ClubDto, len(results))
    for i, club := range results {
        outputs[i] = data.MapClubToClubDto(club)
    }

    return outputs, nil
}