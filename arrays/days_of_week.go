package main

import "fmt"

func main() {
	var intarr = [10]int{
		1, //index=0
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10, //index=9
	}

	var daysOfWeek = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for i := 0; i < len(intarr); i++ {
		fmt.Printf("index=%d element=%d\n", i, intarr[i])
	}

	for index, element := range intarr {
		fmt.Printf("index=%d element=%d\n", index, element)
	}

	for index, day := range daysOfWeek {
		fmt.Printf("index of the day=%d   ", index)
		fmt.Printf("day of the week=%s\n", day)
	}
}
