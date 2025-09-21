package main

import "fmt"

func main() {
	usd := 45
	const fromUSDtoEUR = 0.85
	const fromUSDtoRUB = 83.49
	rstEUR := float64(usd) * fromUSDtoEUR
	rstRUB := float64(usd) * fromUSDtoRUB
	fmt.Println(rstEUR, rstRUB)
}
