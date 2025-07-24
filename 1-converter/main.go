package main

import (
	"fmt"
)

func main() {

	const USD_to_EUR = 0.8554
	const USD_to_RUB = 78.33
	kurs_EUR_to_RUB := USD_to_RUB / USD_to_EUR
	fmt.Print(kurs_EUR_to_RUB)
}
