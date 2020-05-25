package main

import (
	"errors"
	"fmt"
	"os"
)

//LIFO
// [1,2,3,4]
// push 5 [1,2,3,4,5]
// pop (5) [1,2,3,4]

var maxCapacity = 10

func main() {

	var stack []int

	fmt.Println(stack, len(stack), cap(stack))
	stack, err := push(stack, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(stack, len(stack), cap(stack))
	stack, err = push(stack, 2)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(stack, len(stack), cap(stack))

	stack, element, err := pop(stack)
	fmt.Println("popped element is ", element)
	fmt.Println(stack, len(stack), cap(stack))

}

/*
&slice { &pointerToUnderlyingArray, len, cap}
*/
// if there is enough space push the element if not return error
func push(slice []int, element int) ([]int, error) {
	if cap(slice) == maxCapacity {
		return slice, errors.New("no more memory...")
	}

	slice = append(slice, element)
	return slice, nil
}

func pop(slice []int) ([]int, int, error) {
	if len(slice) == 0 {
		return slice, 0, errors.New("no more elements")
	}
	element := slice[len(slice)-1] //copies the elemnts from 0 till length-1
	//arr=[0,1,2,3,4,5]
	//arr[len(arr)-1] = 5
	// arr[0:len(arr-1)] = [0,1,2,3,4]
	slice = slice[0 : len(slice)-1]
	return slice, element, nil
}
