package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

type enrolledCourse struct {
	CourseNumber string
}

func getEnrolledCourse() {
	// Input your student ID here
	studentID := "640612093"

	// URL to scrape
	url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=%s", studentID)

	// Create a new collector
	c := colly.NewCollector()

	// Regular expression to match 6-digit course numbers
	courseNumberRegex := regexp.MustCompile(`\b\d{6}\b`)

	// Slice to store enrolled courses
	var courses []enrolledCourse

	// Set the callback for when an HTML element is found
	c.OnHTML("table[border='0'][cellspacing='1'][cellpadding='3'] tbody tr[bgcolor='#FFFFFF']", func(e *colly.HTMLElement) {
		// Extract the course numbers from the table row
		e.ForEach("td:first-child", func(_ int, td *colly.HTMLElement) {
			// Use regular expression to match 6-digit course numbers
			matches := courseNumberRegex.FindAllString(td.Text, -1)
			// Add each matched course number to the courses slice
			for _, match := range matches {
				courses = append(courses, enrolledCourse{CourseNumber: match})
			}
		})
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	// Print the enrolled courses
	for _, course := range courses {
		fmt.Println(course.CourseNumber)
	}
}