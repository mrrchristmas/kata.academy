package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// DetectOperation find an operation (+ - * /)
func DetectOperation(expression []rune) (operation rune, err error) {
	count := 0
	for _, v := range expression {
		if string(v) == "+" || string(v) == "-" || string(v) == "*" || string(v) == "/" {
			count++
			operation = v
		}
	}
	if count == 1 {
		return operation, nil
	}
	if count == 0 {
		return 0, errors.New("Ошибка, в строке отсутсвует математическая операция")
	}
	if count > 1 {
		return 0, errors.New("Ошибка, в строке слишком много математических операций")
	}
	return operation, nil
}

// FindNumberInString  Get number, numberOfType, err
func FindNumberInString(s string) (int, bool, error) {
	numArab, err := strconv.Atoi(s)
	if err == nil {
		return numArab, true, nil
	}

	numRoman, err := RomanToArab(s)
	if err == nil {
		return numRoman, false, nil
	}

	return 0, false, err
}

// RomanToArab Convert roman to arabic
func RomanToArab(s string) (result int, err error) {
	var RomeAlphabet = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	if strings.Contains(strings.Join(RomeAlphabet, ","), s) == false {
		return 0, errors.New("Не является римским числом")

	}

	Roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}
	lastElement := s[len(s)-1:]
	result = Roman[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if Roman[s[i:i+1]] <= Roman[s[i-1:i]] {
			result += Roman[s[i-1:i]]
		} else {
			result -= Roman[s[i-1:i]]
		}
	}
	return result, nil
}

// Calc get result
func Calc(left, right int, sign int32) (result int, err error) {
	if left > 10 || right > 10 {
		return 0, errors.New("Ошибка, числа больше 10")
	}
	if left < 1 || right < 1 {
		return 0, errors.New("Ошибка, числа меньше 1")
	}
	switch string(sign) {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right != 0 {
			return left / right, nil
		}
	}
	return 0, errors.New("Ошибка, делить на ноль нельзя")
}

// ArabToRoman Convert arabic to roman
func ArabToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}
	return roman
}

func main() {
	fmt.Print("Введите выражение: ")
	in := bufio.NewReader(os.Stdin)
	var s string
	s, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str := []rune(s) //Convert String to []Rune

	//Detect operation
	sign, err := DetectOperation(str) //sign это знак + - * /
	if err != nil {
		panic(err)
	}

	convert := strings.Split(string(str), string(sign)) // Split
	//Find the left number
	leftNumber, leftNumType, err := FindNumberInString(strings.TrimSpace(strings.ToUpper(convert[0])))
	//Find the right number
	rightNumber, rightNumType, err := FindNumberInString(strings.TrimSpace(strings.ToUpper(convert[1])))
	if leftNumber < 1 || leftNumber > 10 || rightNumber < 1 || rightNumber > 10 {
		panic(errors.New("Ошибка, числа должны быть от 1 до 10"))
	}
	//Checking the type of left and right digits
	if leftNumType != rightNumType {
		panic(errors.New("Ошибка, так как используются одновременно разные системы счисления."))
	}
	//Get result Calc
	result, err := Calc(leftNumber, rightNumber, sign)
	if err != nil {
		panic(err)
	}

	if leftNumType == false {
		if result <= 0 {
			panic(errors.New("Ошибка, так как в римской системе нет ноля и отрицательных чисел."))
		}
		fmt.Println(ArabToRoman(result))
		return
	}
	fmt.Println(result)
}
