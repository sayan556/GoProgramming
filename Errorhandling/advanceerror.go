package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	var num float64
	fmt.Print("Enter Number : ")
	fmt.Scanln(&num)
	for i := num; ; i-- {
		value, err := loopExecution(i)
		if err == nil {
			fmt.Println("Result = ", value)
		} else {
			fmt.Println("Error = ", err)
			os.Exit(1)
		}
	}
}
func loopExecution(num float64) (float64, error) {
	if num > 0 {
		return math.Sqrt(num), nil
	} else {
		return 0, errors.New("Math Error : Can not square root of Zero")
	}
}
