package gits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const ID = ""

type Gits struct {
	ID string `json:"id"`
}

var GITS Gits

// func TestPostGit(t *testing.T) {
// 	// Create a new request with GET method and /users endpoint

// 	reqBody := []byte(`{"description":"Example of a gist from golang",
// 		"public":true,
// 		"files":{
// 		"README.md":{
// 			"content":"Hello World"
// 		}
// 	}
// 	}`)
// 	request, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewReader(reqBody))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	token := "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"
// 	request.Header.Set("Authorization", "Bearer "+token)
// 	request.Header.Set("Accept", "application/vnd.github+json")

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Check the status code of the response
// 	if status := response.StatusCode; status != http.StatusCreated {
// 		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
// 	}

// 	defer response.Body.Close()
// 	responseBody, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	//var gits Gits
// 	err = json.Unmarshal(responseBody, &GITS)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(GITS.ID)

// 	// Print the response body
// 	//fmt.Println(string(responseBody))
// }

func TestGetGit(t *testing.T) {
	// Create a new request with GET method and /users endpoint
	request, err := http.NewRequest("GET", "https://api.github.com/gists/b1de2029451d0b3bc2c087afb6b42379", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "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"
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code of the response
	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Print the response body
	fmt.Println(string(responseBody))
}

//func TestUpdateGit(t *testing.T) {
//	// Create a new request with GET method and /users endpoint
//
//	reqBody := []byte(`{
//		"files":{
//		"README.md":{
//			"content":"Hello update"
//		}
//	}
//	}`)
//	request, err := http.NewRequest("PUT", "https://api.github.com/gists", bytes.NewReader(reqBody))
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	token := "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"
//	request.Header.Set("Authorization", "Bearer "+token)
//	request.Header.Set("Accept", "application/vnd.github+json")
//
//	client := &http.Client{}
//	response, err := client.Do(request)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Check the status code of the response
//	if status := response.StatusCode; status != http.StatusOK {
//		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
//	}
//
//	defer response.Body.Close()
//	responseBody, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Print the response body
//	fmt.Println(string(responseBody))
//}
