package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getCPE(year string) {
	token := "8382bd52-4a3d-48ea-8f35-9fc7a3239b7f"
	url := fmt.Sprintf("https://api.cpe.eng.cmu.ac.th/api/v1/curriculum?year=%s&curriculumProgram=CPE&isCOOPPlan=false" , year)

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
    
	fmt.Println(jsonPrettyPrint(string(body)))

}

func jsonPrettyPrint(in string) string {
    var out bytes.Buffer
    err := json.Indent(&out, []byte(in), "", "\t")
    if err != nil {
        return in
    }
    return out.String()
}