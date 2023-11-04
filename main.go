package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
}

func romanToArabicNumeral(roman string) int {
	result := 0
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && romanToArabic[roman[i:i+2]] != 0 {
			result += romanToArabic[roman[i:i+2]]
			i++
		} else {
			result += romanToArabic[string(roman[i])]
		}
	}
	return result
}

func arabicToRomanNumeral(arabic int) string {
	result := ""
	for arabic > 0 {
		if arabic >= 10 {
			result += "X"
			arabic -= 10
		} else if arabic >= 9 {
			result += "IX"
			arabic -= 9
		} else if arabic >= 5 {
			result += "V"
			arabic -= 5
		} else if arabic >= 4 {
			result += "IV"
			arabic -= 4
		} else if arabic >= 1 {
			result += "I"
			arabic -= 1
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный формат ввода")
		return
	}

	a, errA := strconv.Atoi(parts[0])
	if errA != nil {
		a = romanToArabicNumeral(parts[0])
		if a == 0 {
			fmt.Println("Ошибка: неверное число a")
			return
		}
	}

	b, errB := strconv.Atoi(parts[2])
	if errB != nil {
		b = romanToArabicNumeral(parts[2])
		if b == 0 {
			fmt.Println("Ошибка: неверное число b")
			return
		}
	}

	var result int
	switch parts[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неверная операция")
		return
	}

	if errA != nil && errB != nil {
		fmt.Println(arabicToRomanNumeral(result))
	} else {
		fmt.Println(result)
	}
}
