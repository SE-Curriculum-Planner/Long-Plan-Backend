package datanaja

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CPEstudentID struct {
	ID string `json:"ID"`
}

func getCPEstudentID() ([]CPEstudentID, error) {
	var studentIDList []CPEstudentID
	for a := 0; a <= 2; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {

					studentID := fmt.Sprintf("64061%d%d%d%d", a, b, c, d)
					url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data.php?student_id=%s", studentID)
					doc, err := goquery.NewDocument(url)
					if err != nil {
						return nil, err
					}

					// Find and process each block representing courses for a year and semester
					doc.Find("div[class='row marketing']").Each(func(i int, s *goquery.Selection) {
						// Process each row in the table

						s.Find("div[class='col-lg-12'] table[width='100%'][border='0']").Each(func(j int, row *goquery.Selection) {
							if strings.Contains(row.Text(), "วิศวกรรมคอมพิวเตอร์") {
								// Add the course details to the courses slice
								fmt.Print(studentID, " is CPE\n")
								studentIDList = append(studentIDList, CPEstudentID{
									ID: studentID,
								})
							} else {
								fmt.Print(studentID, " is not CPE\n")
							}
						})

					})
				}
			}
		}
	}
	writeCPEStudentIDFile(studentIDList)
	return studentIDList, nil
}

func writeCPEStudentIDFile(studentIDList []CPEstudentID) {
	jsonname := "CPEStudentID.json"
	file, err := os.Create(jsonname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(studentIDList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNumberStudent(student []CPEstudentID) []string {
	var studentIDarr []string
	for _, studentID := range student {
		studentIDarr = append(studentIDarr, studentID.ID)
	}
	return studentIDarr
}

func getDataStudentID(path string) []CPEstudentID {
	jsonFilePath := path

	// Read the JSON file
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
	}

	var response []CPEstudentID
	err = json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
	}

	return response
}
