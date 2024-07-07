package api

import (
	datanaja "github.com/SE-Curriculum-Planner/Long-Plan-Backend/app"
	middlewares "github.com/SE-Curriculum-Planner/Long-Plan-Backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const FIRST_VERSION_PREFIX = "/v1"

func bindFirstVersionRouter(router fiber.Router) {
	firstAPI := router.Group(FIRST_VERSION_PREFIX)
	bindOauthRouter(firstAPI)

	// routes
	firstAPI.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi! welcome to LONGPLAN-API 🌈 \n API Endpoint : \n 1:: /curriculum?major=CPE&year=2563&plan=normal")
	})
	firstAPI.Get("/curriculum", func(c *fiber.Ctx) error {

		major := c.Query("major")
		year := c.Query("year")
		plan := c.Query("plan")
		filename := "./data/curriculum/" + datanaja.GetFilename(major, year, plan)

		// Read JSON file using the function
		jsonFile, err := datanaja.ReadJSONFile(filename)
		if err != nil {
			// Return an error response if unable to read the file
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading JSON file : " + filename)
		}

		// Return curriculum data as JSON response
		return c.JSON(jsonFile)
	})

	firstAPI.Get("/student/enrolledcourses", middlewares.AuthMiddleware(), func(c *fiber.Ctx) error {

		value, ok := c.Locals("STUDENT_CODE").(string)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("studentID not found in context")
		}
		code := value
		
		// Return curriculum data as JSON response
		return c.JSON(datanaja.GetEnrolledCourseByStudentID(code))
	})
}
