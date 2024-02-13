package main

import (
	"fmt"
	"strconv"
)

func main() {
	var aStr, operator, bStr string
	fmt.Println("Введите значение (Пример, 2 + 3, I + V): ")
	fmt.Scanln(&aStr, &operator, &bStr)

	a, err := strconv.Atoi(aStr)
	var isRoman bool
	if err != nil {
		isRoman = true
		a, err = convertToArabic(aStr)
		if err != nil {
			panic(err)
		}
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		isRoman = true
		b, err = convertToArabic(bStr)
		if err != nil {
			panic(err)
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Входные числа должны быть в диапазоне от 1 до 10")
	}

	result, err := calc(a, b, operator)
	if err != nil {
		panic(err)
	}

	if isRoman {
		romanResult, err := convertToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Результат", romanResult)
	} else {
		fmt.Println("Результат", result)
	}
}
func calc(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		} else {
			return a / b, nil
		}
	default:
		return 0, fmt.Errorf("Нет оператора", operator)
	}
}

var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var arabicNumerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func convertToArabic(romanNumeral string) (int, error) {
	number, ok := romanNumerals[romanNumeral]
	if !ok {
		return 0, fmt.Errorf("Недопустимая римская цифра: %s", romanNumeral)
	}
	return number, nil
}

func convertToRoman(arabicNumeral int) (string, error) {
	if arabicNumeral < 1 || arabicNumeral > 10 {
		return "", fmt.Errorf("Недопустимая арабская цифра: %d", arabicNumeral)
	}
	romanNumeral, _ := arabicNumerals[arabicNumeral]
	return romanNumeral, nil
}
