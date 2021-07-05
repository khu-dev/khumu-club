package repository

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/khu-dev/khumu-club/config"
	"github.com/khu-dev/khumu-club/data"
	log "github.com/sirupsen/logrus"
)

var (
	ListUserEndpoint = config.Config.CommandCenterService.Root + "/users"
	ErrInRest        = errors.New("REST API 호출 도중 에러가 발생했습니다 ")
)

func GetUserInfo(token, username string) (*data.User, error) {
	log.Info("khumu-command-center를 이용해 유저 정보를 조회합니다.")
	defer log.Info("khumu-command-center에서 유저 정보를 완료했습니다.")
	resp, err := resty.New().
		R().
		SetAuthToken(token).
		Get(config.Config.CommandCenterService.Root + "/users/me")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if resp.IsError() {
		log.Error(resp.Error())
		return nil, ErrInRest
	}

	userResp := &data.UserResp{}
	err = json.Unmarshal(resp.Body(), userResp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resultUser := data.MapUserDto2User(userResp.Data)
	return resultUser, nil
}

//func IsAdmin(token, username string) (bool, error){
//    resp, err := resty.New().
//        R().
//        SetAuthToken(token).
//        Get(config.Config.CommandCenterService.Root + "/users/me")
//    if err != nil {
//        return false, nil
//    }
//    if resp.IsError() {
//        log.Error(resp.Error())
//        return false, ErrInRest
//    }
//
//    is_admin
//}
