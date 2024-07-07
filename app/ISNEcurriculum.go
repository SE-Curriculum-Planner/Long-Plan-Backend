package datanaja

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

func ISNE_fetch() {
	// Define the mapping
	mapping := map[string]map[string][]string{
		"1": {
			"1": {"001101", "206161", "259104", "259191", "261216", "269101", "269102"},
			"2": {"001102", "140104", "206162", "261205", "269103", "269105", "269130"},
		},
		"2": {
			"1": {"001201", "255201", "261200", "261335", "261336", "269200", "269201", "269202"},
			"2": {"001225", "252281", "261342", "261343", "261433", "269210"},
		},
		"3": {
			"1": {"261305", "261361", "261434", "269340", "269430"},
			"2": {"261446", "261447", "269360", "269370", "269462"},
			"3": {"269401"},
		},
		"4": {
			"1": {"269491"},
			"2": {"259192", "269470", "269492"},
		},
		// Add the rest of the mapping here...
	}

	// Load the existing JSON data
	jsonFile, err := os.Open("data/curriculum/ISNE-2565-normal.json")
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var curriculum Curriculum
	err = json.Unmarshal(byteValue, &curriculum)
	if err != nil {
		return
	}

	// Update the JSON data based on the mapping
	for _, group := range curriculum.CoreAndMajorGroups {
		for i, course := range group.RequiredCourses {
			for year, semesters := range mapping {
				for semester, courseNos := range semesters {
					for _, courseNo := range courseNos {
						if course.CourseNo == courseNo {
							course.RecommendYear, _ = strconv.Atoi(year)
							course.RecommendSemester, _ = strconv.Atoi(semester)
							group.RequiredCourses[i] = course
						}
					}
				}
			}
		}
	}

	// Save the updated JSON data
	file, _ := json.MarshalIndent(curriculum, "", " ")
	_ = ioutil.WriteFile("data/curriculum/ISNE-2565-normal.json", file, 0644)
}
