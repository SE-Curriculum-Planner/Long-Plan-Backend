package datanaja

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CourseData struct {
	CourseNo       string `json:"courseNo"`
	CourseTitleEng string `json:"CourseTitleEng"`
	Abbreviation   string `json:"Abbreviation"`
}

type Course struct {
	CourseNo           string `json:"courseNo"`
	CourseTitleEng     string `json:"courseTitleEng"` // Add this line
	RecommendSemester  int    `json:"recommendSemester"`
	RecommendYear      int    `json:"recommendYear"`
	Prerequisites      []string `json:"prerequisites"`
	Corequisite        *string  `json:"corequisite"`
	Credits            int      `json:"credits"`
}

type CourseGroup struct {
	RequiredCredits int     `json:"requiredCredits"`
	GroupName       string  `json:"groupName"`
	RequiredCourses []Course `json:"requiredCourses"`
	ElectiveCourses []Course `json:"electiveCourses"`
}

type Curriculum struct {
	CurriculumProgram  string        		`json:"curriculumProgram"`
	Year               int           		`json:"year"`
	IsCOOPPlan         bool          		`json:"isCOOPPlan"`
	RequiredCredits    int           		`json:"requiredCredits"`
	FreeElectiveCredits int           		`json:"freeElectiveCredits"`
	CoreAndMajorGroups []CourseGroup 		`json:"coreAndMajorGroups"`
	GeGroups           []CourseGroup     	`json:"geGroups"`
}

type Response struct {
	Ok         bool       `json:"ok"`
	Curriculum Curriculum `json:"curriculum"`
}


func getDataCurriculum(path string) Curriculum {
	jsonFilePath := path

	// Read the JSON file
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
	}

	var response Response
	err = json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
	}

	return response.Curriculum
}

func getCourseNumbersFromCurriculum(curriculum Curriculum) []string {
	var courseNumbers []string

	// Extract from CoreAndMajorGroups
	for _, group := range curriculum.CoreAndMajorGroups {
		courseNumbers = append(courseNumbers, getCourseNumbersFromCourses(group.RequiredCourses)...)
		courseNumbers = append(courseNumbers, getCourseNumbersFromCourses(group.ElectiveCourses)...)
	}

	// Extract from GeGroups
	for _, group := range curriculum.GeGroups {
		courseNumbers = append(courseNumbers, getCourseNumbersFromCourses(group.RequiredCourses)...)
		courseNumbers = append(courseNumbers, getCourseNumbersFromCourses(group.ElectiveCourses)...)
	}

	return courseNumbers
}

func getCourseNumbersFromCourses(courses []Course) []string {
	var courseNumbers []string
	for _, course := range courses {
		courseNumbers = append(courseNumbers, course.CourseNo)
	}
	return courseNumbers
}

func getCourseTitle(courseNumbers []string, curriculum *Curriculum , path string) {
	for _, courseNo := range courseNumbers {
		courseTitles, err := getCourseTitlesFromAPI(courseNo)
		if err != nil {
			fmt.Printf("Error fetching course titles for course number %s: %v\n", courseNo, err)
			continue
		}

		for _, courseTitle := range courseTitles {
	
			// Now, you can use the 'apiURL' as needed.
			fmt.Printf("Course Number: %s, Course Title: %s\n", courseNo, courseTitle)
			// Update the curriculum with the course title

			updateCurriculumWithCourseTitleAndWriteToFile(curriculum , courseNo, courseTitle , path)
		}
	}

}

func updateCurriculumWithCourseTitleAndWriteToFile(curriculum *Curriculum, courseNo string, courseTitle string , outputPath string) error {
	// Helper function to update course titles in a slice of Course
	updateCourseTitles := func(courses []Course) {
		for i, course := range courses {
			if course.CourseNo == courseNo {
				courses[i].CourseTitleEng = courseTitle
			}
		}
	}

	// Update course titles in CoreAndMajorGroups
	for i := range curriculum.CoreAndMajorGroups {
		updateCourseTitles(curriculum.CoreAndMajorGroups[i].RequiredCourses)
		updateCourseTitles(curriculum.CoreAndMajorGroups[i].ElectiveCourses)
	}

	// Update course titles in GeGroups
	for i := range curriculum.GeGroups {
		updateCourseTitles(curriculum.GeGroups[i].RequiredCourses)
		updateCourseTitles(curriculum.GeGroups[i].ElectiveCourses)
	}

	// Write the updated curriculum back to a file
	return writeCurriculumToFile(curriculum, outputPath)
}

func writeCurriculumToFile(curriculum *Curriculum, outputPath string) error {
	// Convert curriculum to JSON
	curriculumJSON, err := json.MarshalIndent(curriculum, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling curriculum to JSON: %v", err)
	}

	// Write JSON to file
	err = ioutil.WriteFile(outputPath, curriculumJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing curriculum JSON to file: %v", err)
	}

	return nil
}

func getCourseTitlesFromAPI(courseNo string) ([]string, error) {
	apiURL := fmt.Sprintf("https://mis-api.cmu.ac.th/tqf/v1/course-template?courseid=%s&academicyear=2563&academicterm=1", courseNo)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var courses []CourseData
	err = json.Unmarshal(body, &courses)
	if err != nil {
		return nil, err
	}

	var courseTitles []string
	for _, course := range courses {
		if course.Abbreviation != "" {
			courseTitles = append(courseTitles, course.Abbreviation) // if the course has an abbreviation
		} else {
			courseTitles = append(courseTitles, course.CourseTitleEng) // if the course doesn't have an abbreviation, use the long name
		}
	}
	return courseTitles, nil
}


