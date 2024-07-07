package datanaja

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type enrolledCourse struct {
	Year         string `json:"year"`
	Semester     string `json:"semester"`
	CourseNumber string `json:"courseNo"`
	Credit       string `json:"credit"`
	Grade        string `json:"grade"`
}

func getEnrolledCourses(studentID string) ([]enrolledCourse, error) {
	// Configure custom HTTP transport with TLS settings.
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false}, // WARNING: InsecureSkipVerify should be false in production.
	}

	// Create an HTTP client with the custom transport and a timeout.
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second, // Set a timeout to prevent indefinite hanging.
	}

	// Construct the URL with the student ID.
	url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=%s", studentID)

	// Use the custom client to make the request.
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	// Ensure the response status is OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the document from the response body.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %v", err)
	}
	var courses []enrolledCourse
	var year, semester string

	// Find and process each block representing courses for a year and semester
	doc.Find("table[width='100%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		// Process each row in the table
		s.Find("td[align='center'] > B").Each(func(j int, row *goquery.Selection) {
			// Check if the text contains Thai characters for semester and year
			if strings.Contains(row.Text(), "ภาคเรียนที่") && strings.Contains(row.Text(), "ปีการศึกษา") {
				// Extract semester and year from the text
				semesterMatches := regexp.MustCompile(`\d+`).FindAllString(row.Text(), -1)
				if len(semesterMatches) >= 2 {
					semester = semesterMatches[len(semesterMatches)-2]
					temp_year := semesterMatches[len(semesterMatches)-1]

					num_year, nil := strconv.Atoi(temp_year)
					if err != nil {
						fmt.Println("Error converting year to int:", err)
						return
					}
					num_studentID, nil := strconv.Atoi(transformInput(studentID))
					if err != nil {
						fmt.Println("Error converting student to int:", err)
						return
					}

					num_year = num_year - num_studentID + 1 // calculated study year by studentID and academic year
					year = strconv.Itoa(num_year)

				}
			}

		})
		s.Find("table[cellspacing='1'][cellpadding='3'][width='60%'][border='0'][class='t'] tr[bgcolor='#FFFFFF']").Each(func(j int, row *goquery.Selection) {
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

	})

	return courses, nil
}

func groupCoursesByYearSemester(courses []enrolledCourse) map[string]map[string][]enrolledCourse {
	groupedCourses := make(map[string]map[string][]enrolledCourse)

	for _, course := range courses {
		if _, ok := groupedCourses[course.Year]; !ok {
			groupedCourses[course.Year] = make(map[string][]enrolledCourse)
		}
		groupedCourses[course.Year][course.Semester] = append(groupedCourses[course.Year][course.Semester], course)
	}

	// Ensure at least 4 years of data
	for year := 1; year <= 4; year++ {
		yearStr := strconv.Itoa(year)
		if _, ok := groupedCourses[yearStr]; !ok {
			groupedCourses[yearStr] = make(map[string][]enrolledCourse)
		}
		// Ensure there are two semesters for each year
		if _, ok := groupedCourses[yearStr]["1"]; !ok {
			groupedCourses[yearStr]["1"] = []enrolledCourse{}
		}
		if _, ok := groupedCourses[yearStr]["2"]; !ok {
			groupedCourses[yearStr]["2"] = []enrolledCourse{}
		}
	}

	return groupedCourses
}

func writeGroupedToFile(groupedCourses map[string]map[string][]enrolledCourse, fileName string) {

	if len(groupedCourses) == 0 {
		fmt.Println("No data to write. Grouped courses map is empty.")
		return
	}

	jsonname := fmt.Sprintf("data/student-courseEnrolled/" + fileName + "-grouped-enrolled.json")
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

func encodeGroupedToJSON(groupedCourses map[string]map[string][]enrolledCourse) string {
	if len(groupedCourses) == 0 {
		return "Course Not Found"
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ") // Set indentation for a pretty format
	err := encoder.Encode(groupedCourses)
	if err != nil {
		return "Error to Encode"
	}

	// Convert byte slice to string
	jsonString := buffer.String()

	// Return JSON data as string
	return jsonString
}

func writeToFile(courses []enrolledCourse, fileName string) {
	jsonname := fmt.Sprintf("data/student-courseEnrolled/" + fileName + "-enrolled.json")
	file, err := os.Create(jsonname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(courses)
	if err != nil {
		log.Fatal(err)
	}
}
