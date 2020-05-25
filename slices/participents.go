package main

import "fmt"

func main() {
	var slice = []int{}
	fmt.Println(slice, len(slice), cap(slice))

	var slice2 = []int{1, 2, 3}
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Printf("memory address=%p\n", slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Printf("memory address=%p\n", slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Printf("memory address=%p\n", slice2)
}
