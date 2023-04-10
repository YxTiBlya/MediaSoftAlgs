package main

import "fmt"

func main() {
	var num string
	fmt.Scanln(&num)

	result := 0
	for i, value := range num {
		switch value {
		case 'I':
			result += 1
		case 'V':
			result += 5

			if i > 0 {
				if num[i-1] == 'I' {
					result -= 2
				}
			}
		case 'X':
			result += 10

			if i > 0 {
				if num[i-1] == 'I' {
					result -= 2
				}
			}
		case 'L':
			result += 50

			if i > 0 {
				if num[i-1] == 'X' {
					result -= 20
				}
			}
		case 'C':
			result += 100

			if i > 0 {
				if num[i-1] == 'X' {
					result -= 20
				}
			}
		case 'D':
			result += 500

			if i > 0 {
				if num[i-1] == 'C' {
					result -= 200
				}
			}
		case 'M':
			result += 1000

			if i > 0 {
				if num[i-1] == 'C' {
					result -= 200
				}
			}
		}
	}

	fmt.Println(result)
}
