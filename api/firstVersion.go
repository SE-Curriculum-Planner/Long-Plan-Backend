package api

import "github.com/gofiber/fiber/v2"

const FIRST_VERSION_PREFIX = "/v1"

func bindFirstVersionRouter(router fiber.Router) {
	firstAPI := router.Group(FIRST_VERSION_PREFIX)

	bindOauthRouter(firstAPI)

}
