package main

import (
	"fmt"
)

func main() {
	count_valut := UserInput()
	const USD_to_EUR float64 = 0.8554
	const USD_to_RUB float64 = 78.33
	var summ_new_valute float64
	kurs_EUR_to_RUB := USD_to_RUB / USD_to_EUR
	summ_new_valute = count_valut * USD_to_RUB / USD_to_EUR
	fmt.Println(kurs_EUR_to_RUB)
	fmt.Print(summ_new_valute)
}

func UserInput() float64 {
	var count_valut float64
	fmt.Print("Введите количество исходной валюты: ")
	fmt.Scan(&count_valut)
	return count_valut
}
func calculate(count_valut float64, start_valut int, final_valut int) {

}
