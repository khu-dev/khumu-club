package data

import "github.com/khu-dev/khumu-club/ent"

// 유저의 요청에서의 입력용으로도 사용하고
// 유저의 요청에 대한 응답용으로도 사용
type ClubDto struct{
    Name *string `json:"name"`
    Summary *string `json:"summary"`
    Description *string `json:"description"`
    Hashtags []string `json:"hashtags"`
    Images []string `json:"images"`
    Homepage *string `json:"homepage"`
    Instagram *string `json:"instagram"`
    Facebook *string `json:"facebook"`
    Phone *string `json:"phone"`
    Email *string `json:"email"`
    Recommended *bool `json:"recommended"`
    Liked *bool `json:"liked"`
    LikeCount int `json:"like_count"`
}

func MapClubToClubDto(source *ent.Club) *ClubDto {
    output := &ClubDto{
        Name: &source.Name,
        Summary: &source.Summary,
        Description: &source.Description,
        Hashtags: source.Hashtags,
        Images: source.Images,
        Homepage: String2StringPtr(source.Homepage, true),
        Instagram: String2StringPtr(source.Instagram, true),
        Facebook: String2StringPtr(source.Facebook, true),
        Phone: String2StringPtr(source.Phone, true),
        Email: String2StringPtr(source.Email, true),
        Recommended: &source.Recommended,
    }

    return output
}