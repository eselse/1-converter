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
