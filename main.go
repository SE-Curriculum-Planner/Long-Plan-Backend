package main

import (
	"fmt"
	"log"
)

func main() {
	// getCPEstudentID()
	// countCPEstudent()

	fetchCPEcurriculumAndMap()

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


func mapCPEcourseToCMUapi(path string) {
	responseNormal := getDataCurriculum(path)
	courseNumbersNormal := getCourseNumbersFromCurriculum(responseNormal)
	totalMembers := len(courseNumbersNormal)
	fmt.Println("Course count : " , totalMembers)
	getCourseTitle(courseNumbersNormal , &responseNormal , path)
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