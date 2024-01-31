package main

import (
	"fmt"
	"log"
)

func main() {
	// getCPEAPI("2563" , "CPE" , "true")
	// getCPEAPI("2563" , "CPE" , "false")
	
	// mapCPEcourseToCMUapi()
	getEnrolledCourseByStudentID("640612093")
	// getAllCurriculumYear("CPE")
	// getAllCurriculumYear("ISNE")

	// getCourses() // old method
}

func getAllCurriculumYear(major string) {
	for i := 58; i <= 67; i++ {
		year := fmt.Sprintf("25%d",i)
		getCPEAPI(year,major,"true")
		getCPEAPI(year,major,"false")
	}
}

func mapCPEcourseToCMUapi() {
	responseNormal := getData("data/curriculum/CPE-2563-normal.json")
	courseNumbersNormal := getCourseNumbersFromCurriculum(responseNormal.Curriculum)
	totalMembers := len(courseNumbersNormal)
	fmt.Println("Course count : " , totalMembers)
	getCourseTitle(courseNumbersNormal)
}

func getEnrolledCourseByStudentID(id string) {
	// Input your student ID here
	studentID := id

	courses, err := getEnrolledCourses(studentID)
	if err != nil {
		log.Fatal(err)
	}
	groupedCourses := groupCoursesByYearSemester(courses)
	writeGroupedToFile(groupedCourses, studentID)
	// writeToFile(courses,studentID)
}
