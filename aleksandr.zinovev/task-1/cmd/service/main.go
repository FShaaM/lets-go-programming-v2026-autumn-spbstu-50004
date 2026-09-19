package main

import (
	"fmt"
)

func main() {
	var operand1, operand2 int
	var operation string

	_, err1 := fmt.Scan(&operand1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err2 := fmt.Scan(&operand2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err3 := fmt.Scan(&operation)
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(operand1 + operand2)
	case "-":
		fmt.Println(operand1 - operand2)
	case "*":
		fmt.Println(operand1 * operand2)
	case "/":
		if operand2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(operand1 / operand2)
	default:
		fmt.Println("Invalid operation")
	}
}
