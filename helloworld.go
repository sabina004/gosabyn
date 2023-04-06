// Write a program that takes a list of integers as input and returns a new list containing only the odd numbers from the original list.
package main

import "fmt"

func OddNumbers(numbers []int) []int {
	var oddNumbers []int
	for _, num := range numbers {
		if num%2 != 0 {
			oddNumbers = append(oddNumbers, num)
		}
	}
	return oddNumbers
}

func main() {
	var num int
	numbers := []int{}
	for {
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}
	oddNumbers := OddNumbers(numbers)
	fmt.Println("Original List: ", numbers)
	fmt.Println("Odd Numbers List: ", oddNumbers)
}
