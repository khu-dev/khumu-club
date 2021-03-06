package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khu-dev/khumu-club/adapter/slack"
	"github.com/khu-dev/khumu-club/data"
	"github.com/khu-dev/khumu-club/repository"
	"github.com/khu-dev/khumu-club/service"
	"github.com/khu-dev/khumu-club/util"
	log "github.com/sirupsen/logrus"
)

type ClubHandler struct {
	ClubService  *service.ClubService
	SlackAdapter slack.SlackAdapter
}

func (h *ClubHandler) CreateClub(c *fiber.Ctx) error {
	user := util.GetRequestUser(c)
	token := util.GetRequestToken(c)
	user, err := repository.GetUserInfo(token, user.ID)
	if err != nil {
		log.Error("유저 정보를 가져오지 못했습니다.")
		return fiber.ErrUnauthorized
	}

	isAdmin := service.IsAdmin(user)
	if !isAdmin {
		log.Error("관리자가 아닌 유저의 요청입니다.")
		return fiber.ErrForbidden
	}

	c.Context().Request.Header.Header()
	body := &data.ClubDto{}
	err = c.BodyParser(body)
	if err != nil {
		log.Error(err)
		return c.Status(500).SendString(err.Error())
	}

	club, err := h.ClubService.CreateClub(body)
	if err != nil {
		log.Error(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"data":       club,
		"message":    "",
		"error_type": "",
	})
}

// UserGet returns a user
func (h *ClubHandler) ListClubs(c *fiber.Ctx) error {
	category := c.Query("category", "")
	log.Info("category: ", category)
	if len(category) != 0 {
		clubs, err := h.ClubService.ListClubsByCategoryContaining(category)
		if err != nil {
			log.Error(err)
		}
		return c.JSON(fiber.Map{
			"data":       clubs,
			"message":    err,
			"error_type": "",
		})
	}
	clubs, err := h.ClubService.ListClubs()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"data":       clubs,
		"message":    "",
		"error_type": "",
	})
}

func (h *ClubHandler) ClubAddRequest(c *fiber.Ctx) error {
	body := &data.ClubAddOrModifyRequestDto{}
	err := c.BodyParser(body)
	if err != nil {
		log.Error(err)
		return err
	}

	if err = h.ClubService.ClubAddRequest(c, body); err != nil {
		log.Error(err)
		return err
	}

	log.Info("동아리 추가 요청을 받았습니다.", body)

	return c.JSON(fiber.Map{
		"data":       nil,
		"message":    "동아리 추가를 요청했습니다.",
		"error_type": "",
	})
}

func (h *ClubHandler) ClubModifyRequest(c *fiber.Ctx) error {
	body := &data.ClubAddOrModifyRequestDto{}
	err := c.BodyParser(body)
	if err != nil {
		log.Error(err)
		return err
	}

	if err = h.ClubService.ClubAddRequest(c, body); err != nil {
		log.Error(err)
		return err
	}

	log.Info("동아리 수정 요청을 받았습니다.", body)

	return c.JSON(fiber.Map{
		"data":       nil,
		"message":    "동아리 수정을 요청했습니다.",
		"error_type": "",
	})
}

func (h *ClubHandler) ListCategories(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"data":       h.ClubService.ListCategories(c),
		"message":    nil,
		"error_type": nil,
	})
}

//// UserCreate registers a user
//func UserCreate(c *fiber.Ctx) error {
//	user := &models.User{
//		Name: c.FormValue("user"),
//	}
//	database.Insert(user)
//	return c.JSON(fiber.Map{
//		"success": true,
//		"user":    user,
//	})
//}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendString("404 NotFound")
}
