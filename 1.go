package main

import (
	"fmt"
	"strconv"
)

func main() {
	for num := 1; num <= 100; num++ {
		str := ""
		if num%3 == 0 {
			str += "Fizz"
		}
		if num%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = strconv.Itoa(num)
		}
		fmt.Println(str)
	}
}
