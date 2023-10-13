package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func modulus(a float64) float64 {
	// The modulus operation for a single float64 value can be calculated using math.Mod function.
	// The second argument can be 2 to get the remainder after division by 2.
	return math.Mod(a, 2)
}

func squareRoot(a float64) float64 {
	return math.Sqrt(a)
}

func reciprocal(a float64) float64 {
	if a == 0 {
		return 0
	}
	return 1 / a
}

func factorial(a float64) float64 {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}

func sine(angle float64) float64 {
	return math.Sin(angle)
}

func cosine(angle float64) float64 {
	return math.Cos(angle)
}

func tangent(angle float64) float64 {
	return math.Tan(angle)
}

func inputs() (float64, float64) {
	var num1, num2 float64

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	return num1, num2
}

func inputsinglenum() float64 {
	var num1 float64

	fmt.Print("Enter the number: ")
	fmt.Scan(&num1)

	return num1
}

func angle() float64 {
	var num1 float64

	fmt.Print("Enter the angle: ")
	fmt.Scan(&num1)
	return num1 * (math.Pi / 180.0)
}

func main() {
	var operator, condition string
	condition = "Y"
	for condition != "N" {

		fmt.Print("Enter the operation (+, -, *, /,%,sqrt,reciprocal,factorial,sine,cosine,tangent): ")
		fmt.Scan(&operator)

		var result float64
		var err error

		switch operator {
		case "+":
			result = add(inputs())
		case "-":
			result = subtract(inputs())
		case "*":
			result = multiply(inputs())
		case "/":
			result, err = divide(inputs())
		case "%":
			result = modulus(inputsinglenum())
		case "sqrt":
			result = squareRoot(inputsinglenum())
		case "reciprocal":
			result = reciprocal(inputsinglenum())
		case "factorial":
			result = factorial(inputsinglenum())
		case "sine":
			result = sine(angle())
		case "cosine":
			result = cosine(angle())
		case "tangent":
			result = tangent(angle())
		default:
			fmt.Println("Invalid operation.")
			return
		}

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.2f\n", result)
		}
		fmt.Println("Do you want to continue(Y/N)?")
		fmt.Scan(&condition)
	}
}
