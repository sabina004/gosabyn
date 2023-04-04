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
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	oddNumbers := OddNumbers(numbers)
	fmt.Println("Original List: ", numbers)
	fmt.Println("Odd Numbers List: ", oddNumbers)
}
