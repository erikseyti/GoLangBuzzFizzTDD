package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := fizzBuzz(50)

	fmt.Println(result)

}

func fizzBuzz (i int) string {
	if i %3 ==0 && i%5 == 0 {
		return "fizzbuzz"
	} else if i %3 == 0 {
		return "fizz"
	} else if i % 5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(i)
}
