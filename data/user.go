package data

// 내부적으로 사용하기 위함
type User struct{
    ID string `json:"id"`
    Groups []string `json:"groups"`
}

// 내부적으로 사용하기 위함
type Group struct{
    ID int `json:"id"`
    Name string `json:"name"`
}

// 유저 서비스에 API 요청에 대한 응답에 사용되는 User 타입
type UserDto struct{
    Username string `json:"username"`
    Groups []*Group `json:"groups"`
}

// 유저 서비스의 API 요청 형태
type UserResp struct{
    Message *string `json:"message"`
    Data *UserDto `json:"data"`
}

func MapUserDto2User(src *UserDto) *User{
    dest := &User{}
    dest.ID = src.Username
    dest.Groups = make([]string, len(src.Groups))
    for i, group := range src.Groups {
        dest.Groups[i] = group.Name
    }

    return dest
}