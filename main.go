package main

import (
	"fmt"
)

func main() {
	// 1. Concatenate strings and store the result in a variable
	var concatenatedString string
	string1 := "Hello"
	string2 := "World"
	concatenatedString = string1 + " " + string2

	// 2. Use fmt.Sprintf(...) to format strings using %s and %d, and print the value
	formattedString := fmt.Sprintf("The concatenated string is: %s, and the length is: %d", concatenatedString, len(concatenatedString))
	fmt.Println(formattedString)

	// 3. Perform mathematical operations: +, -, *, / with any numbers, store the results in variables, and print them
	var additionResult, subtractionResult, multiplicationResult, divisionResult float64
	num1 := 10.5
	num2 := 5.2
	additionResult = num1 + num2
	subtractionResult = num1 - num2
	multiplicationResult = num1 * num2
	divisionResult = num1 / num2
	fmt.Println("Addition Result:", additionResult)
	fmt.Println("Subtraction Result:", subtractionResult)
	fmt.Println("Multiplication Result:", multiplicationResult)
	fmt.Println("Division Result:", divisionResult)

	// 4. Create two float variables with different values and write four if conditional statements using operators: >, >=, <, <=, ==, !=, &&, ||. Also, add else to two conditions, and print the operations inside the conditions
	num3 := 8.7
	num4 := 3.4

	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else {
		fmt.Println("num1 is not greater than num2")
	}

	if num1 >= num2 {
		fmt.Println("num1 is greater than or equal to num2")
	} else {
		fmt.Println("num1 is not greater than or equal to num2")
	}

	if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 is not less than num2")
	}

	if num1 <= num2 {
		fmt.Println("num1 is less than or equal to num2")
	} else {
		fmt.Println("num1 is not less than or equal to num2")
	}

	if num1 == num2 {
		fmt.Println("num1 is equal to num2")
	} else {
		fmt.Println("num1 is not equal to num2")
	}

	if num1 != num2 {
		fmt.Println("num1 is not equal to num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	if num3 > num2 && num3 < num4 {
		fmt.Println("num3 is greater than num2 and less than num4")
	}

	if num3 > num2 || num3 < num4 {
		fmt.Println("num3 is greater than num2 or less than num4")
	}

	// 5. Create a string variable and print its value using switch case with at least three cases
	var conditionString = "second"

	switch conditionString {
	case "first":
		fmt.Println("This is the first case")
	case "second":
		fmt.Println("This is the second case")
	case "third":
		fmt.Println("This is the third case")
	default:
		fmt.Println("This is the default case")
	}

	// 6. Create another switch case without a switch parameter and print phrases from any two case statements
	switch {
	case num1 > num2:
		fmt.Println("num1 is greater than num2")
		fallthrough
	case num1 < num2:
		fmt.Println("num1 is less than num2")
	}

	// 7. Create a switch case with a default statement and print something from the default block
	switch conditionString {
	case "first":
		fmt.Println("This is the first case")
	default:
		fmt.Println("This is the default case")
	}

	// 8. Create an array of numbers with a length of 5, increment the element at index 3, decrement the element at index 4, and print the result
	array := [5]int{1, 2, 3, 4, 5}
	array[3]++
	array[4]--
	fmt.Println(array)

	// 9. Create two string slices in two different ways (with and without using make), add a string from point 1 to the first slice, print it, also print the number of elements in the second slice
	slice1 := make([]string, 0)
	slice2 := []string{"apple", "banana", "orange"}

	concatenatedString = string1 + " " + string2
	slice1 = append(slice1, concatenatedString)
	fmt.Println(slice1)

	fmt.Println("Number of elements in slice2:", len(slice2))
}
