package http

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/khu-dev/khumu-club/data"
	"github.com/khu-dev/khumu-club/repository"
	"github.com/khu-dev/khumu-club/service"
	log "github.com/sirupsen/logrus"
)

type ClubHandler struct{
	ClubService *service.ClubService
}

func GetRequestUser(c *fiber.Ctx) *data.User{
	a := c.Locals("user")
	log.Println(a)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id, ok := claims["user_id"].(string)
	if !ok {
		log.Error("토큰 내의 해당 Claim을 찾을 수 없습니다.")
		return nil
	}

	return &data.User{
		ID: id,
	}
}

func GetRequestToken(c *fiber.Ctx) string{
	user := c.Locals("user").(*jwt.Token)
	return user.Raw
}

func (h *ClubHandler) CreateClub(c *fiber.Ctx) error {
	user := GetRequestUser(c)
	token := GetRequestToken(c)
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
		"data": club,
		"message":"",
		"error_type":"",
	})
}

// UserGet returns a user
func (h *ClubHandler) ListClub(c *fiber.Ctx) error {
	clubs, err := h.ClubService.ListClub()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"data": clubs,
		"message":"",
		"error_type":"",
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
