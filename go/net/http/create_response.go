package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// JSON string
	d := "{\"test\": \"Hello world\"}"
	resp := &http.Response{
		Header:     http.Header{},
		StatusCode: http.StatusSeeOther,
		Body:       ioutil.NopCloser(bytes.NewBufferString(d)),
	}
	resp.Header.Set("Location", "location/url")
	fmt.Printf("resp = %+v\n", resp)
	defer resp.Body.Close()

	// Define struct for JSON.
	var body struct {
		Test string `json:"test"`
	}

	// Decode JSON from stream.
	json.NewDecoder(resp.Body).Decode(&body)
	fmt.Printf("body = %+v\n", body)
	// Ignore Upper/Lower
	fmt.Printf("location = %+v\n", resp.Header.Get("location"))
}
