package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getData(path string) Response{
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

	return response	
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

func getCourseTitle(courseNumbers []string) {
	for _, courseNo := range courseNumbers {
		courseTitles, err := getCourseTitlesFromAPI(courseNo)
		if err != nil {
			fmt.Printf("Error fetching course titles for course number %s: %v\n", courseNo, err)
			continue
		}

		for _, courseTitle := range courseTitles {
			// Construct the API URL using courseNo and courseTitle

			// Now, you can use the 'apiURL' as needed.
			fmt.Printf("Course Number: %s, Course Title: %s\n", courseNo, courseTitle)
		}
	}
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
			courseTitles = append(courseTitles, course.Abbreviation)  // if course have abbreviation
		} else {
			courseTitles = append(courseTitles, course.CourseTitleEng) // if course don't have abbreviation then put long name into it
		}
		
	}

	return courseTitles, nil
}