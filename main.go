package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type ApiResponse struct {
	// Define fields in the struct that match the expected JSON response structure
	// Adjust these fields based on the actual response structure from the API
	TQF2CopyToEbulletinID  string `json:"TQF2CopyToEbulletinID"`
	AcademicTerm           int    `json:"AcademicTerm"`
	AcademicYear           string `json:"AcademicYear"`
	CurriculumCentralCode  string `json:"CurriculumCentralCode"`
	CurriculumNameEng      string `json:"CurriculumNameEng"`
	CurriculumNameTha      string `json:"CurriculumNameTha"`
	DegreeFullEng          string `json:"DegreeFullEng"`
	DegreeShortEng         string `json:"DegreeShortEng"`
	DepartmentID           string `json:"DepartmentID"`
	FacultyID              string `json:"FacultyID"`
	IsNewCurriculum        bool   `json:"IsNewCurriculum"`
	StudentLevelID         string `json:"StudentLevelID"`
	SysCreateDate          string `json:"SysCreateDate"`
	SysUpdateDate          string `json:"SysUpdateDate"`
	TQF2BranchID           string `json:"TQF2BranchID"`
	FacultyBranchNameTha   string `json:"FacultyBranchNameTha"`
	TQF2CurriculumTypeID   int    `json:"TQF2CurriculumTypeID"`
	VersionNumber          int    `json:"VersionNumber"`
	StudentLevelNameTha    string `json:"StudentLevelNameTha"`
	StudentLevelNameEng    string `json:"StudentLevelNameEng"`
	IsClose                bool   `json:"IsClose"`
	CloseTime              string `json:"CloseTime"`
	CloseDate              string `json:"CloseDate"`
}

var curriculum_ID string

func main() {
	studentID := "640612093"
	studentFaculty := "Computer Engineering"
	output := transformInput(studentID)
	fmt.Printf("Student ID : %s \n" , studentID)
	fmt.Printf("Student curriculum year : %s \n" , output)
	fmt.Printf("Student Faculty : %s \n" , studentFaculty)
	apiURL := fmt.Sprintf("https://mis-api.cmu.ac.th/tqf/v1/tqf2/copy-to-ebulletin/ebulletin-public?studentlevelid=2&facultyid=06&academicYear=%s&acdemicterm=1", output)

    resp, err := http.Get(apiURL)
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    // Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status:", resp.Status)
		return
	}

	// Decode the JSON response into an array of ApiResponse structs
	var apiResponses []ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponses)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	var FID string
	// Iterate through the array of responses and print DegreeShortEng for matching studentFaculty
	for _, apiResponse := range apiResponses {
		if strings.Contains(apiResponse.CurriculumNameEng, studentFaculty) {
			fmt.Printf("Faculty ID: %s\n", apiResponse.TQF2CopyToEbulletinID)
			FID = apiResponse.TQF2CopyToEbulletinID
		}
	}
	curriculum_ID = "https://mis.cmu.ac.th/TQF/TQF2/CurriculumPublic.aspx?EID=" + FID
	
	courses, err := scraping(curriculum_ID)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	for _, course := range courses {
		fmt.Printf("Course ID: %s\nCourse Code: %s\nCourse Short Code: %s\nCourse Title: %s\n\n",
			course.ID, course.CourseCode, course.CourseShortCode, course.CourseTitle)
	}
}

func transformInput(input string) string {
	// Define a regular expression pattern to extract the first two digits.
	re := regexp.MustCompile(`^(\d{2})`)

	// Find the first two digits in the input.
	matches := re.FindStringSubmatch(input)

	if len(matches) != 2 {
		// If the pattern is not found, return an error or handle the case accordingly.
		return "Invalid input"
	}

	// Extract the first two digits and concatenate with "25".
	firstTwoDigits := matches[1]
	result :=  "25" + firstTwoDigits

	// Convert the result to an integer and then back to a string to remove leading zeros.
	resultInt, _ := strconv.Atoi(result)
	result = strconv.Itoa(resultInt)

	return result
}