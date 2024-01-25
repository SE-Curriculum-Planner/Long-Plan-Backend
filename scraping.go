package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Course struct {
	ID                 string
	CourseCode         string
	CourseShortCode    string
	CourseTitle		   string
	CourseType 		   string 
}

func scraping(curriculum_ID string) ([]Course, error) {
	var courses []Course
	fmt.Printf("Start Scarping URL : %v\n", curriculum_ID)
	// Make the HTTP request
	response, err := http.Get(curriculum_ID)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	// Extract information using dynamic IDs
	document.Find("div[id^='GVCurriculumList_ctl'][id$='_DvCourse']").Each(func(index int, divHtml *goquery.Selection) {
		// Extract course details within the div
		courseID := strings.TrimSpace(divHtml.Find("[id$='_lblCourseID']").Text())
		courseCode := strings.TrimSpace(divHtml.Find("[id$='_lblCourseCode']").Text())
		courseShortCode := strings.TrimSpace(divHtml.Find("[id$='_lblCourseShort']").Text())
		courseTitle := strings.TrimSpace(divHtml.Find("[id$='_lblCourseTitle']").Text())

		// Append the course to the courses slice
		courses = append(courses, Course{
			ID:              courseID,
			CourseCode:      courseCode,
			CourseShortCode: courseShortCode,
			CourseTitle:     courseTitle,
		})
	})
	return courses, nil
}


