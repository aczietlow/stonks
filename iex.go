package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type api struct {
	baseurl  string
	version  string
	apiToken string
}

// hardcode the hell out of this
func quote(token string, ticker string) string {
	type stockData struct {
		Symbol        string
		LastSalePrice float32
	}

	iexApi := api{
		baseurl:  "https://cloud.iexapis.com",
		version:  "stable",
		apiToken: token,
	}
	endpoint := iexApi.baseurl + "/" + iexApi.version + "/tops/?token=" + token + "&symbols=" + ticker
	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("API Error:", err)
		log.Fatal("Failed to connect.")
	}

	defer resp.Body.Close()

	// Deal the error handling.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data []stockData
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	if len(data) < 1 {
		return fmt.Sprintf("Unable to find any information on %v are you sure the ticker is correct?", ticker)
	}
	return fmt.Sprintf("%.2f", data[0].LastSalePrice)
}
