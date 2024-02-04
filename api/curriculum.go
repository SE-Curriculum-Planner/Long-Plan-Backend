package api

import (
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

const CURRICULUM_PREFIX = "/curriculum"

func bindCurriculumRouter(router fiber.Router) {
	curriculum := router.Group(CURRICULUM_PREFIX)

	hdl := handler.NewCurriculumHandler()
}
