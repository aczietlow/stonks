package main

import (
	"log"
	"net/http"
)

// Gets dog images
func getCat() string {
	resp, err := http.Get("https://thecatapi.com/api/images/get")
	// http client follows the redirect, no need to track the 301 response.
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp.Request.URL.String()
}
