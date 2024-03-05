package datanaja

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getCPEAPI(year string , major string , coop string) {
	token := "8382bd52-4a3d-48ea-8f35-9fc7a3239b7f"
	url := fmt.Sprintf("https://api.cpe.eng.cmu.ac.th/api/v1/curriculum?year=%s&curriculumProgram=%s&isCOOPPlan=%s" , year , major , coop)

    // Create a Bearer string by appending string access token
    var bearer = "Bearer " + token

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error while GET : %s",url)
	}
    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }

    	
	var data map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Check if the "ok" field is true
	okValue, ok := data["ok"].(bool)
	if !ok || !okValue {
        if(coop == "true") {
		    fmt.Printf("%s-%s Curriculum COOP-PLAN is not found!!! \n" , major , year)
        } else {
            fmt.Printf("%s-%s Curriculum NORMAL-PLAN is not found!!! \n" , major , year)
        }
		return
	}

    writeFile(coop,major,year,string(body))

    fmt.Printf("===============================================\n")
    if coop == "true" {
        fmt.Printf("%s-%s COOP-PLAN is found!!!\n",major,year)
	    fmt.Printf("JSON data exported to %s-%s-coop.json\n",major,year)
    } else {
        fmt.Printf("%s-%s NORMAL-PLAN is found!!!\n",major,year)
	    fmt.Printf("JSON data exported to %s-%s-normal.json\n",major,year)
    }
    fmt.Printf("===============================================\n")
}

func jsonPrettyPrint(in string) string {
    var out bytes.Buffer
    err := json.Indent(&out, []byte(in), "", " ")
    if err != nil {
        return in
    }
    return out.String()
}

func writeFile(coop string , major string , year string , body string) {
    // Open a file for writing (create if not exists, truncate if exists)
    if coop == "true" {
        filename := fmt.Sprintf("data/curriculum/"+"%s-%s-coop.json", major ,year)
	    file, err := os.Create(filename)
	    if err != nil {
		fmt.Println("Error creating file:", err)
		return
	    }
	    defer file.Close()

	    // Write the original JSON string to the file
	    _, err = file.WriteString(jsonPrettyPrint(string(body)))
	    if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	    }
    } else {
        filename := fmt.Sprintf("data/curriculum/"+"%s-%s-normal.json", major ,year)
	    file, err := os.Create(filename)
	    if err != nil {
		fmt.Println("Error creating file:", err)
		return
	    }
	    defer file.Close()

	    // Write the original JSON string to the file
	    _, err = file.WriteString(jsonPrettyPrint(string(body)))
	    if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	    }
    }
    
}