package main

import "frontendmasters.com/go/crypto/api"

func main() {
	rate, err := api.GetRate("btc")
	print(rate, err)
}
