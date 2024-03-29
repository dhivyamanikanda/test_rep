package gits

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const ID = ""

type Gits struct {
	ID string `json:"id"`
}

type Config struct {
	Git string `json:"git"`
}

var GITS Gits

//func TestPostGit(t *testing.T) {
//	// Create a new request with GET method and /users endpoint
//
//	reqBody := []byte(`{"description":"Example of a gist from golang",
//		"public":true,
//		"files":{
//		"README.md":{
//			"content":"Hello World"
//		}
//	}
//	}`)
//	request, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewReader(reqBody))
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
//	if status := response.StatusCode; status != http.StatusCreated {
//		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
//	}
//
//	defer response.Body.Close()
//	responseBody, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	//var gits Gits
//	err = json.Unmarshal(responseBody, &GITS)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Println(GITS.ID)
//
//	// Print the response body
//	//fmt.Println(string(responseBody))
//}

func TestGetGit(t *testing.T) {
	// file, err := os.Open("env.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // Parse the JSON into a Config struct
	// var config Config
	// err = json.NewDecoder(file).Decode(&config)
	// if err != nil {
	// 	panic(err)
	// }

	// Set the environment variables
	// err = os.Setenv("config.git", config.Git)
	// if err != nil {
	// 	t.Errorf("error set env %v", err)
	// }
	// fmt.Println("config.git")
	// fmt.Println(config.Git)

	// Create a new request with GET method and /users endpoint
	request, err := http.NewRequest("GET", "https://api.github.com/gists/5d7e231294d51571a661eb8b823c379b", nil)
	if err != nil {
		t.Fatal(err)
	}

	// err = os.Getenv("config.git", config.Git)
	
	// fmt.Println(os.Getenv("vars.TOKEN"))
	//token := "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"
	request.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN"))
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
