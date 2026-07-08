package main

import "fmt"

type ConversionResult struct {
	Currency string
	Amount   float64
}

var rates = map[string]float64{
	"USD": 15000,
	"EUR": 16000,
	"JPY": 140,
	"SGD": 11000,
}

func main() {
	var conversion []ConversionResult
	var input float64
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&input)
	// for i := 0; i < 4; i++ {
	// 	conversion = append(conversion, ConversionResult{Currency: rates[i], Amount: rates[i] * input})
	// }

	for currency, rate := range rates {
		conversion = append(conversion, ConversionResult{Currency: currency, Amount: input / rate})
	}

	fmt.Println("Konversi dari $", input)
	// fmt.Println("USD", int64(input/rates["USD"]))
	// fmt.Println("EUR", int64(input/rates["EUR"]))
	// fmt.Println("JPY", int64(input/rates["JPY"]))
	// fmt.Println("SGD", int64(input/rates["SGD"]))

	for _, amount := range conversion {
		fmt.Printf("%s: %d\n", amount.Currency, int64(amount.Amount))
	}

}
