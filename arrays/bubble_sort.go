package main

import "fmt"

func main() {
	//var arr = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	//arr := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	arr := [...]int{
		10,
		9,
		8,
		7,
		6,
		5,
		4,
		3,
		2,
		1,
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	fmt.Printf("capcity=%d length=%d\n", cap(arr), len(arr))

	fmt.Println(arr)
}
