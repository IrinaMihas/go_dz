package main

import (
	"fmt"
)

const USD_to_EUR float64 = 0.8554
const USD_to_RUB float64 = 78.33

func main() {
	valut1 := UserInput_nachValut()
	fmt.Println(valut1)
	v_count := Vvod_count_valut()
	fmt.Println(v_count)
	valut2 := UserInput_endValut(valut1)
	fmt.Println(valut2)
	kurs := calculate(v_count, valut1, valut2)
	fmt.Println("После обмена ", valut1, " на ", valut2, "вы получите: ", kurs)

}

func calculate(count_valut float64, start_valut string, final_valut string) float64 {
	var summ_new_valut float64
	kurs_EUR_to_RUB := USD_to_RUB / USD_to_EUR
	switch {
	case start_valut == "EUR":
		if final_valut == "RUB" {
			summ_new_valut = count_valut * USD_to_RUB / USD_to_EUR
		} else {
			summ_new_valut = count_valut * USD_to_EUR
		}
		return summ_new_valut
	case start_valut == "RUB":
		if final_valut == "EUR" {
			summ_new_valut = count_valut / kurs_EUR_to_RUB
			return summ_new_valut
		} else {
			summ_new_valut = count_valut / USD_to_RUB
		}

	case start_valut == "USD":
		if final_valut == "EUR" {
			summ_new_valut = count_valut * USD_to_EUR
		} else {
			summ_new_valut = count_valut * USD_to_RUB
		}

	}
	return summ_new_valut
}

func UserInput_nachValut() string {
	var valuta string
	for {
		fmt.Print("Введите исходную валюту(RUB, EUR или USD): ")
		fmt.Scan(&valuta)

		if valuta == "EUR" || valuta == "RUB" || valuta == "USD" {
			break
		} else {
			fmt.Println("Неверный ввод! Пожалуйста, введите USD, RUB или EUR")
		}
	}

	return valuta
}

func Vvod_count_valut() float64 {
	var v_count float64
	for {
		fmt.Print("Введите количество валюты(больше нуля): ")
		fmt.Scan(&v_count)

		if v_count > 0 {
			break
		} else {
			fmt.Println("Неверный ввод! Пожалуйста, введите значение больше нуля")
		}
	}

	return v_count
}

func UserInput_endValut(nachvalut string) string {
	var valuta string
	for {
		fmt.Print("Введите целевую валюту: ")
		fmt.Scan(&valuta)

		if valuta == "EUR" && valuta != nachvalut || valuta == "RUB" && valuta != nachvalut || valuta == "USD" && valuta != nachvalut {
			break
		} else {
			fmt.Println("Неверный ввод! Пожалуйста, введите USD, RUB или EUR")
		}
	}

	return valuta
}
