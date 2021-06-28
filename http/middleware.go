package http

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v2"
    log "github.com/sirupsen/logrus"
    "os"
)

func NewJWTMiddleware() fiber.Handler{
    return jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("KHUMU_SECRET")),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			log.Info("requestUser=", GetRequestUser(ctx))
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				authorizationHeader := ctx.Request().Header.Peek("Authorization")
				if len(authorizationHeader) == 0 {
					log.Info("미인증 요청입니다.")
					return ctx.Next()
				} else {
					return ctx.Status(fiber.StatusBadRequest).SendString("Missing or malformed JWT")
				}
			} else {
				return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
			}
		},
	})
}

