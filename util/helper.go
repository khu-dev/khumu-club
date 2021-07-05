package util

import (
    "github.com/form3tech-oss/jwt-go"
    "github.com/gofiber/fiber/v2"
    "github.com/khu-dev/khumu-club/data"
    log "github.com/sirupsen/logrus"
)

func GetRequestUser(c *fiber.Ctx) *data.User{
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
