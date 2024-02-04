package handler

import (
	"log"
	"time"

	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/config"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/errors"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/lodash"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/pkg/oauth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/samber/lo"
)

type oauthHandler struct {
}

func NewOauthHandler() *oauthHandler {
	return &oauthHandler{}
}

func (h oauthHandler) SignIn(c *fiber.Ctx) error {
	config := config.Config.Application
	code := c.Query("code", "")
	if lo.IsEmpty(code) {
		return lodash.ResponseBadRequest(c)
	}
	user, err := oauth.CmuOauthValidation(code)
	if err != nil {
		return lodash.ResponseError(c, errors.NewStatusBadGatewayError(err.Error()))
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &oauth.UserClaims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * 24 * time.Hour)),
		},
	})

	token, err := claims.SignedString([]byte(config.Secret))
	if err != nil {
		log.Print(err)
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.MaxAge = 5 * 24 * 3600
	cookie.Path = "/"
	cookie.HTTPOnly = true
	cookie.SameSite = "lax"
	cookie.Domain = config.Domain
	c.Cookie(cookie)
	return lodash.ResponseOK(c, user.Cmuitaccount)
}

func (h oauthHandler) GetUser(c *fiber.Ctx) error {
	token := c.Cookies("token")
	config := config.Config.Application
	parsedAccessToken, err := jwt.ParseWithClaims(token, &oauth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}
	user := &parsedAccessToken.Claims.(*oauth.UserClaims).User
	return lodash.ResponseOK(c, user)
}

func (h oauthHandler) Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return lodash.ResponseNoContent(c, nil)
}
