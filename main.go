package main

import "fmt"

func main() {
	usd := 45
	const fromUSDtoEUR = 0.85
	const fromEURtoRUB = 83.49
	rstEUR := float64(usd) * fromUSDtoEUR
	rstRUB := rstEUR * fromEURtoRUB
	fmt.Println(rstRUB)
}

func getUserInput() float64 {
	usd := 0.0
	fmt.Println("Enter amount in USD:")
	fmt.Scanln(&usd)
	return usd
}

func convertUSDtoEUR(amount, usd, eur float64) float64 { return 0 }
