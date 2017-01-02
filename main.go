package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type stockQuote struct {
	Status           string  `json:"Status"`
	Name             string  `json:"Name"`
	Symbol           string  `json:"Symbol"`
	LastPrice        float64 `json:"LastPrice"`
	Change           float64 `json:"Change"`
	ChangePercent    float64 `json:"ChangePercent"`
	Timestamp        string  `json:"Timestamp"`
	MSDate           float64 `json:"MSDate"`
	MarketCap        float64 `json:"MarketCap"`
	Volume           int32   `json:"Volume"`
	ChangeYTD        float64 `json:"ChangeYTD"`
	ChangePercentYTD float64 `json:"ChangePercentYTD"`
	High             float64 `json:"High"`
	Low              float64 `json:"Low"`
	Open             float64 `json:"Open"`
}

func main() {
	symbol := "MSFT"

	// QueryEscape escape the string so it can be safely placed inside a URL query
	safeSymbol := url.QueryEscape(symbol)

	url := fmt.Sprintf("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=%s", safeSymbol)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Create a HTTP client
	client := &http.Client{}

	// Do sends a HTTP request and returns a HTTP response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Callers should close the resp.Body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record stockQuote

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		panic(err)
	}

	fmt.Println("Company Name: ", record.Name)
	fmt.Println("Last Value: ", record.LastPrice)

}
