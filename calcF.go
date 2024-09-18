package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var romanNumerals = map[rune]int{
	'I': 1, 'V': 5, 'X': 10,
}

func romanToInt(s string) (int, error) {
	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue, exists := romanNumerals[rune(s[i])]
		if !exists {
			return 0, fmt.Errorf("ПАНИКА")
		}
		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}

	if total < 1 || total > 10 {
		return 0, fmt.Errorf("ПАНИКА")
	}

	return total, nil
}

func isRoman(s string) bool {
	for _, char := range s {
		if _, exists := romanNumerals[char]; !exists {
			return false
		}
	}
	return true
}

func intToRoman(num int) string {
	romanVal := []struct {
		val int
		sym string
	}{
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	result := ""
	for _, v := range romanVal {
		for num >= v.val {
			result += v.sym
			num -= v.val
		}
	}
	return result
}

func isValidArabic(num int) bool {
	return num >= 1 && num <= 10
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Что считаем?")

	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	re := regexp.MustCompile(`^\s*([\d]+|[IVX]+)\s*([+\-*\/])\s*([\d]+|[IVX]+)\s*$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		panic("ПАНИКА")
	}

	aStr, operator, bStr := matches[1], matches[2], matches[3]

	var a, b int
	var err error

	if isRoman(aStr) && isRoman(bStr) {
		a, err = romanToInt(aStr)
		if err != nil {
			panic(err)
		}
		b, err = romanToInt(bStr)
		if err != nil {
			panic(err)
		}
	} else {
		a, err = strconv.Atoi(aStr)
		if err != nil || !isValidArabic(a) {
			panic("ПАНИКА!!")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil || !isValidArabic(b) {
			panic("ПАНИКА!!")
		}
	}

	var result int

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("ПАНИКА!")
		}
		result = a / b
	default:
		panic("ПАНИКА!")
	}

	if isRoman(aStr) {
		if result < 1 {
			panic("ПАНИКА!!")
		}
		fmt.Printf("Результат: %s\n", intToRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}