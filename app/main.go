package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// getCPEstudentID()
	// countCPEstudent()

	// getEnrolledCourseByStudentID("640612093")
	apiTEST()
	
	// getEnrolledCourseByStudentID("640612093")
	// fetchCPEcurriculumAndMap()

	// getAllCurriculumYear("CPE")
	// getAllCurriculumYear("ISNE")

	// getCoursesByStudentIDandFaculty("640612093" , "Computer Engineering")

	// getCourses() // old method
}

func getAllCurriculumYear(major string) {
	for i := 58; i <= 67; i++ {
		year := fmt.Sprintf("25%d",i)
		getCPEAPI(year,major,"true")
		getCPEAPI(year,major,"false")
	}
}

func apiTEST() {
	
	// fiber instance
	app := fiber.New()

	app.Use(cors.New())
   
	// routes 
	app.Get("/", func(c *fiber.Ctx) error {
	 return c.SendString("Hi! welcome to LONGPLAN-API 🌈 \n API Endpoint : \n 1:: /curriculum?major=CPE&year=2563&plan=normal \n 2:: /student/enrolledcourses?studentID={input}")
	})
	app.Get("/curriculum", func(c *fiber.Ctx) error {
		
		major := c.Query("major")
		year := c.Query("year")
		plan := c.Query("plan")
		filename := "data/curriculum/" + getFilename(major,year,plan)

		// Read JSON file using the function
		jsonFile, err := readJSONFile(filename)
		if err != nil {
			// Return an error response if unable to read the file
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading JSON file : " + filename)
		}
	
		// Return curriculum data as JSON response
		return c.JSON(jsonFile)
	   })
	
	   app.Get("/student/enrolledcourses", func(c *fiber.Ctx) error {
		
		studentID := c.Query("studentID")
		
		// Return curriculum data as JSON response
		return c.JSON(getEnrolledCourseByStudentID(studentID))
	   })

   // app listening at PORT: 3000
	app.Listen(":3000")
}

func readJSONFile(filePath string) (*Curriculum, error) {
    // Read the JSON file
    jsonFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        // Return error if unable to read the file
        return nil, err
    }

    // Initialize a Curriculum struct to hold the parsed JSON data
    var curriculum Curriculum

    // Unmarshal the JSON data into the Curriculum struct
    err = json.Unmarshal(jsonFile, &curriculum)
    if err != nil {
        // Return error if unable to unmarshal JSON
        return nil, err
    }

    // Return the parsed Curriculum struct
    return &curriculum, nil
}

func getFilename(major, year, plan string) string {
    // Generate filename based on parameters
    // Example: CPE-2023-normal.json
    return strings.ToUpper(major) + "-" + year + "-" + plan + ".json"
}

func mapCPEcourseToCMUapi(path string) {
	responseNormal := getDataCurriculum(path)
	courseNumbersNormal := getCourseNumbersFromCurriculum(responseNormal)
	totalMembers := len(courseNumbersNormal)
	fmt.Println("Course count : " , totalMembers)
	getCourseTitle(courseNumbersNormal , &responseNormal , path)
}

func getEnrolledCourseByStudentID(id string) map[string]map[string][]enrolledCourse {
	// Input your student ID here
	studentID := id

	courses, err := getEnrolledCourses(studentID)
	if err != nil {
		log.Fatal(err)
	}
	groupedCourses := groupCoursesByYearSemester(courses)

	return groupedCourses
	// writeToFile(courses,studentID)
}

func fetchCPEcurriculumAndMap() {
	getAllCurriculumYear("CPE")
	normal := "data/curriculum/CPE-2563-normal.json"
	coop := "data/curriculum/CPE-2563-coop.json"
	
	mapCPEcourseToCMUapi(normal)
	mapCPEcourseToCMUapi(coop)
}

func countCPEstudent() {
	responseNormal := getDataStudentID("data/student-courseEnrolled/CPEStudentID.json")
	CPEStudent := getNumberStudent(responseNormal)
	totalMembers := len(CPEStudent)
	fmt.Println("Number of CPE Student : " , totalMembers)
}
