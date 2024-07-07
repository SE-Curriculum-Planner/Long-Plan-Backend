package datanaja

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type scrapCourse struct {
	ID              string
	CourseCode      string
	CourseShortCode string
	CourseTitle     string
	Credit          string
	CourseType      string
}

func scraping(curriculum_ID string) ([]scrapCourse, error) {
	var courses []scrapCourse
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

	// Extract information within tbody > tr
	document.Find("div[id^='GVCurriculumList_ctl'][id$='_DvCourse']").Each(func(index int, trHtml *goquery.Selection) {
		// Extract course details within the tr
		courseID := strings.TrimSpace(trHtml.Find("[id$='_lblCourseID']").Text())
		courseCode := strings.TrimSpace(trHtml.Find("[id$='_lblCourseCode']").Text())
		courseShortCode := strings.TrimSpace(trHtml.Find("[id$='_lblCourseShort']").Text())
		courseTitle := strings.TrimSpace(trHtml.Find("[id$='_lblCourseTitle']").Text())
		courseType := strings.TrimSpace(trHtml.Find("[id$='_lblCurriculumMainStructureName']").Text())
		courseCredit := strings.TrimSpace(trHtml.Find("[id$='_lblCreditShow']").Text())

		// Append the course to the courses slice
		courses = append(courses, scrapCourse{
			ID:              courseID,
			CourseCode:      courseCode,
			CourseShortCode: courseShortCode,
			CourseTitle:     courseTitle,
			CourseType:      courseType,
			Credit:          courseCredit,
		})
	})

	return courses, nil
}
