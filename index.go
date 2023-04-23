package main

import "fmt"

func main() {

	var result string
	var checkNumber int = 10
	if isOdd := isNumberIsOdd(checkNumber); isOdd {
		result = "odd number"
	} else {
		result = "even number"
	}

	fmt.Printf("\n%d is %s", checkNumber, result)

	const a, b = -1, -15
	fmt.Printf("\nmin of %d and %d is %d", a, b, findMin(a, b))

	const getEnumCategory = "PROGRAMMING"
	fmt.Printf("\nEnum of category %s is %d", getEnumCategory, getEnumOfCategory(getEnumCategory))

}

func isNumberIsOdd(number int) bool {
	return number%2 == 1
}

func findMin(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func getEnumOfCategory(category string) int {
	switch category {
	case "MATH":
		return 1
	case "SCIENCE":
		return 2
	case "LANGUAGE":
		return 3
	case "PROGRAMMING":
		return 4
	default:
		return 0
	}
}
