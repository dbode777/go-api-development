package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) int {
	return a / b
}

func divisionUsingErrorsPackage(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New(": cannot divide by zero")
	}
	return a / b, nil
}

// Use defer to recover from a panic. It waits to execute until the end of the function.
func main() {

	fmt.Println(divisionUsingErrorsPackage(10, 0))

	defer func() {
		// recover is a built-in function that recovers from a panic and only works inside a deferred function. This essentially mimics a try/catch block or try/except block.
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		// can also use panic to crash the program if that is expected
	}()
	fmt.Println(divide(10, 0))

}
