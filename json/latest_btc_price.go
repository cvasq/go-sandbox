package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Currency struct {
	Last float64 `json:"last"`
}

type CurrencyList struct {
	USD Currency `json:"USD"`
}

func main() {

	resp, err := http.Get("https://blockchain.info/ticker")
	if err != nil {
		log.Println(err)
	}

	currentPrice := CurrencyList{}
	err = json.NewDecoder(resp.Body).Decode(&currentPrice)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	lastPrice := fmt.Sprintf("$%.2f", currentPrice.USD.Last)

	fmt.Println(lastPrice)
}
