package main

import "fmt"

func main() {
	getAllCurriculumYear("CPE")
	getAllCurriculumYear("ISNE")
	// getCourses() // old method
}

func getAllCurriculumYear(major string) {
	for i := 58; i <= 67; i++ {
		year := fmt.Sprintf("25%d",i)
		getCPEAPI(year,major)
	}
}