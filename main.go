package main

import (
	"fmt"
)

type CourseData struct {
	CourseNo       string  `json:"courseNo"`
	CourseTitleEng string  `json:"CourseTitleEng"`
	Abbreviation   string  `json:"Abbreviation"`
}

type Course struct {
	CourseNo           string   `json:"courseNo"`
	RecommendSemester  int      `json:"recommendSemester"`
	RecommendYear      int      `json:"recommendYear"`
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

type GeGroup struct {
	RequiredCredits int           `json:"requiredCredits"`
	GroupName       string        `json:"groupName"`
	RequiredCourses []Course      `json:"requiredCourses"`
	ElectiveCourses []Course      `json:"electiveCourses"`
}

type Curriculum struct {
	CurriculumProgram string         `json:"curriculumProgram"`
	Year              int            `json:"year"`
	IsCOOPPlan        bool           `json:"isCOOPPlan"`
	RequiredCredits   int            `json:"requiredCredits"`
	FreeElectiveCredits int          `json:"freeElectiveCredits"`
	CoreAndMajorGroups []CourseGroup  `json:"coreAndMajorGroups"`
	GeGroups           []GeGroup      `json:"geGroups"`
}

type Response struct {
	Ok         bool       `json:"ok"`
	Curriculum Curriculum `json:"curriculum"`
}

func main() {
	// getCPEAPI("2563" , "CPE" , "true")
	// getCPEAPI("2563" , "CPE" , "false")
	
	// responseNormal := getData("CPE-2563-normal.json")
	// courseNumbersNormal := getCourseNumbersFromCurriculum(responseNormal.Curriculum)
	// totalMembers := len(courseNumbersNormal)
	// fmt.Println("Course count : " , totalMembers)
	// getCourseTitle(courseNumbersNormal)

	getEnrolledCourse()

	// getData("CPE-2563-normal.json")
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

