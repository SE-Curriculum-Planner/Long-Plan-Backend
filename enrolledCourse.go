package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type enrolledCourse struct {
	Year         int    `json:"Year"`
	Semester     int    `json:"Semester"`
	CourseNumber string `json:"CourseNumber"`
	Credit       string `json:"Credit"`
	Grade        string `json:"Grade"`
}

func getEnrolledCourses(studentID string) ([]enrolledCourse, error) {
	url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=%s", studentID)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var courses []enrolledCourse
	year, semester := 1, 1

	// Find and process each block representing courses for a year and semester
	doc.Find("table[cellspacing='1'][cellpadding='3'][width='60%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		// Process each row in the table
		s.Find("tr[bgcolor='#FFFFFF']").Each(func(j int, row *goquery.Selection) {
			// Extract course details from each row
			courseNumber := strings.TrimSpace(row.Find("td:first-child").Text())
			credit := strings.TrimSpace(row.Find("td:nth-child(2)").Text())
			grade := strings.TrimSpace(row.Find("td:nth-child(3)").Text())

			// Add the course details to the courses slice
			courses = append(courses, enrolledCourse{
				CourseNumber: courseNumber,
				Credit:       credit,
				Grade:        grade,
				Semester:     semester,
				Year:         year,
			})
		})

		// Increment semester and year appropriately
		semester++
		if semester > 2 {
			semester = 1
			year++
		}
	})

	return courses, nil
}

func groupCoursesByYearSemester(courses []enrolledCourse) map[int]map[int][]enrolledCourse {
	groupedCourses := make(map[int]map[int][]enrolledCourse)

	for _, course := range courses {
		if _, ok := groupedCourses[course.Year]; !ok {
			groupedCourses[course.Year] = make(map[int][]enrolledCourse)
		}
		groupedCourses[course.Year][course.Semester] = append(groupedCourses[course.Year][course.Semester], course)
	}

	return groupedCourses
}

func writeToFile(courses []enrolledCourse, fileName string) {
	jsonname := fmt.Sprintf(fileName+"-enrolled.json")
	file, err := os.Create(jsonname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Set indentation for a pretty format
	err = encoder.Encode(courses)
	if err != nil {
		log.Fatal(err)
	}
}

func writeGroupedToFile(groupedCourses map[int]map[int][]enrolledCourse, fileName string) {
	jsonname := fmt.Sprintf(fileName+"-grouped-enrolled.json")
	file, err := os.Create(jsonname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Set indentation for a pretty format
	err = encoder.Encode(groupedCourses)
	if err != nil {
		log.Fatal(err)
	}
}
