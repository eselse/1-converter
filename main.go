package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Currency Converter")
	result := convertCurrency(getUserInput())
	fmt.Printf("Converted amount: %.2f\n", result)
}

func getCurrency() string {
	currency := ""
	for {
		fmt.Scanln(&currency)
		if currency != "USD" && currency != "EUR" && currency != "RUB" {
			fmt.Println("Invalid currency. Please enter USD, EUR, or RUB.")
			continue
		} else {
			break
		}
	}
	return currency
}

func getFloat() float64 {
	for {
		reader := bufio.NewReader(os.Stdin)
		inputFloat, _ := reader.ReadString('\n')
		inputFloat = strings.TrimSpace(inputFloat)

		floatNum, errFloat := strconv.ParseFloat(inputFloat, 64) // 64 specifies float64

		if errFloat != nil {
			fmt.Printf("Error: '%s' is not a valid decimal number.\nEnter amount.\n", inputFloat)
			continue
		} else if floatNum < 0 {
			fmt.Println("Error: Amount cannot be negative. Please enter a positive number.")
			continue
		} else {
			return floatNum
		}
	}
}

func getUserInput() (float64, string, string) {
	fmt.Println("Enter a first currency (USD, EUR, RUB): ")
	firstCurrency := getCurrency()
	fmt.Println("Enter amount: ")
	amount := getFloat()
	fmt.Println("Enter a first currency (USD, EUR, RUB): ")
	secondCurrency := getCurrency()
	return amount, firstCurrency, secondCurrency
}

func convertCurrency(amount float64, firstCurrency, secondCurrency string) float64 {
	result := 0.0
	if firstCurrency == "USD" && secondCurrency == "EUR" {
		result = amount * 0.92
	} else if firstCurrency == "EUR" && secondCurrency == "USD" {
		result = amount * 1.09
	} else if firstCurrency == "USD" && secondCurrency == "RUB" {
		result = amount * 82.23
	} else if firstCurrency == "RUB" && secondCurrency == "USD" {
		result = amount * 0.012
	} else if firstCurrency == "EUR" && secondCurrency == "RUB" {
		result = amount * 89.32
	} else if firstCurrency == "RUB" && secondCurrency == "EUR" {
		result = amount * 0.011
	} else if firstCurrency == secondCurrency {
		result = amount
	}

	return result
}
