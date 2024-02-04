package api

import (
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/internal/handler"
	middlewares "github.com/SE-Curriculum-Planner/Long-Plan-Backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const OAUTH_PREFIX = "/oauth"

func bindOauthRouter(router fiber.Router) {
	oauth := router.Group(OAUTH_PREFIX)

	hdl := handler.NewOauthHandler()

	oauth.Post("", hdl.SignIn)
	oauth.Get("/me", middlewares.AuthMiddleware(), hdl.GetUser)
	oauth.Post("/signout", middlewares.AuthMiddleware(), hdl.Logout)
}
