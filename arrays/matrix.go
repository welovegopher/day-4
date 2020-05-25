package main

import (
	"fmt"
	"strings"
)

//indentity matrix
/*
1, 0, 0
0, 1, 0
0, 0, 1
*/
func main() {
	matrix := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	//fmt.Println(matrix)

	//for rowIndex, row := range matrix {
	//	fmt.Println(rowIndex, row)
	//}

	for _, row := range matrix {
		fmt.Println(row)
		fmt.Println(row[1])
		fmt.Println(strings.Repeat("*", 20))
	}
}
