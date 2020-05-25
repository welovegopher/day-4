package main

import "fmt"

//in go [10]int{} is not same as []int{}
func main() {
	arr := [10]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	printEvenNumbers(arr)
	printOddNumbers(arr)

	PrintWithSomeCondition(arr, isEven)
	PrintWithSomeCondition(arr, isOdd)

	//find if a given element is present in array or not

	PrintWithSomeCondition(arr, linearsearch)
	PrintWithSomeCondition(arr, func(i int) bool {
		searchElement := 5
		if searchElement == i {
			return true
		}
		return false
	})

}

func linearsearch(i int) bool {
	searchElement := 5
	if searchElement == i {
		return true
	}
	return false
}

//DRY (do not repeat yourself)

func PrintWithSomeCondition(arr [10]int, condition func(int) bool) {
	for i, element := range arr {
		//instead of hardcoding the condition lets say even or odd ... we ask the caller to provide the logic
		if condition(element) {
			fmt.Printf("condition met at index=%d value=%d\n", i, element)
		}
	}
}

func printEvenNumbers(arr [10]int) {
	for i, element := range arr {
		if isEven(element) { //if element is even
			fmt.Printf("even number found at index=%d and element=%d\n", i, element)
		}
	}
}

func isEven(i int) bool {
	return i%2 == 0
}

func isOdd(i int) bool {
	return i%2 != 0
}

func printOddNumbers(arr [10]int) {
	for i, element := range arr {
		if isOdd(element) { //if element is odd
			fmt.Printf("odd number found at index=%d and element=%d\n", i, element)
		}
	}
}
