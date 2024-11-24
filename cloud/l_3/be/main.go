package main

import "fmt"

func Add(a int, b int) int {

	return a + b
}

func Multiply(a int, b int) int {

	return a * b
}

func Divide(a int, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {

	fmt.Println("Addition:", Add(10, 5))            // 15
	fmt.Println("Multiplication:", Multiply(10, 5)) // 50

	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result) // 5
	}
}
