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

	fmt.Printf("%d is %s", checkNumber, result)

}

func isNumberIsOdd(number int) bool {
	return number%2 == 1
}
