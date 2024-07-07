package middlewares

import (
	"log"

	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/config"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/errors"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/lodash"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/oauth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/samber/lo"
)

func AuthMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		config := config.Config.Application

		invalidToken := errors.NewUnauthorizedError(errors.AuthErr("invalid token").Error())

		token := c.Cookies("token")
		if lo.IsEmpty(token) {
			return lodash.ResponseError(c, errors.NewUnauthorizedError("empty token"))
		}

		parsedAccessToken, err := jwt.ParseWithClaims(token, &oauth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err != nil {
			log.Print(err)
			return lodash.ResponseError(c, invalidToken)
		}
		user := &parsedAccessToken.Claims.(*oauth.UserClaims).User

		c.Locals("ROLE", user.ItaccounttypeEN)
		c.Locals("STUDENT_CODE", user.StudentID)
		return c.Next()
	}
}
