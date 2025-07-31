package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	var err error
	var strNumbers []string

	for {
		fmt.Print("Введите операцию вычисления (AVG, SUM, MED): ")

		operation, err = scanOperation()
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите числа через запятую: ")
		strNumbers = scanNumbers()
		result := operateNumbers(strNumbers, operation)
		fmt.Printf("Результат вычисления %.v = %.v\n", operation, result)
		break
	}

}

func scanOperation() (string, error) {
	var operation string
	fmt.Scan(&operation)

	if operation == "AVG" || operation == "SUM" || operation == "MED" {
		return operation, nil
	} else {
		err := errors.New("некорректное значение операции")
		return "", err
	}

}

func scanNumbers() []string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
	}

	numbers := strings.Split(input, ",")

	return numbers
}

func operateNumbers(strNumbers []string, operation string) float64 {
	floatNumbers := []float64{}
	result := 0.0

	for _, strNum := range strNumbers {
		floatNum, _ := strconv.ParseFloat(strNum, 64)

		floatNumbers = append(floatNumbers, floatNum)
	}

	switch operation {
	case "AVG":
		sum := 0.0
		for _, num := range floatNumbers {
			sum += num
		}
		result = sum / float64(len(floatNumbers))
	case "SUM":
		sum := 0.0
		for _, num := range floatNumbers {
			sum += num
		}
		result = sum
	case "MED":
		sort.Float64s(floatNumbers)
		len := len(floatNumbers)

		if len%2 == 0 {
			result = (floatNumbers[len/2-1] + floatNumbers[len/2]) / 2
		}
		result = floatNumbers[len/2]

	}

	return result
}
