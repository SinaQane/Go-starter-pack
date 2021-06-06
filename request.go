package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type IpifyResult struct {
	IpAddress string `json:"ip"` // This is used for decoding a json file into an object of this struct
}

// Default timeout time for http.Client is infinite, so we'll create a GlobalHttpClient for the entire project

var GlobalHttpClient = &http.Client{Timeout: time.Second * 5}

func main() {

	// Making a new request with arbitrary method and doing it later

	request, _ := http.NewRequest("DELETE", "https://postman-echo.com/delete", nil)
	_, _ = GlobalHttpClient.Do(request) // Returns response and error

	// Using the response...
}

func normalGet() {

	// Sending a request to an URL and getting the result back

	resp, err := http.Get("http://api.ipify.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode, resp.Status) // 404 NOT FOUND, 403 FORBIDDEN, 200 OK
	body, err := io.ReadAll(resp.Body) // Only supported in Go 1.16 and higher
	_ = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func readJsonNormal() {

	/* Reading a json result from a request and decoding it into an object using json.Unmarshal
	json.Unmarshal takes a []byte and converts it into json.
	 */

	resp, err := http.Get("http://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode, resp.Status)
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	var result IpifyResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.IpAddress)
}

func readJsonStream() {

	// Same as globalJsonStream(), but this one isn't using GlobalHttpClient. So its runtime limit is infinity

	resp, err := http.Get("http://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	var result IpifyResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	_ = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.IpAddress)
}

func globalJsonStream() {

	/* Reading a json result from a request and decoding it into an object using json.NewDecoder
	json.NewDecoder takes a IO Streamer (Reader) and converts it into json.
	*/

	resp, err := GlobalHttpClient.Get("http://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	var result IpifyResult
	err = json.NewDecoder(resp.Body).Decode(&result) // Decoding resp.Body into &result
	_ = resp.Body.Close() // Again, resp.Body should be closed right away
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.IpAddress)
}

func postToURL() {

	/* For converting an object into json, it's better to use "sprintf" and hard-coding the json,
	but "sprintf" gets a bunch of errors while working with strings.
	e.g. "HELL\nO", etc. So we'll use json.Marshal function:
	 */

	jsonMarshalled, _ := json.Marshal(IpifyResult{IpAddress: "192.168.1.1"}) // Converting an object into json byte slice
	resp, err := GlobalHttpClient.Post("https://postman-echo.com/post", "application/json", bytes.NewReader(jsonMarshalled))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close() // Always close the resp.Body after reading it (the "log.fatal()" in the next line may cause the program to close and memory leak can occur)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}