package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Region  string `json:"region"`
}

func main() {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("name", "Germany").
		Get("https://restcountries.com/v3.1/name")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var countries []Country
	err = json.Unmarshal(resp.Body(), &countries)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(countries) == 0 {
		fmt.Println("No countries found")
		return
	}

	country := countries[0]
	fmt.Printf("%s: %s\n", country.Name, country.Capital)
}
