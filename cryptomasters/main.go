package main

import (
	"fmt"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	rate, err := api.GetRate("btc")
	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}
