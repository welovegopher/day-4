package main

import (
	"fmt"
	"strings"
)

//calculate no. of times an element is repeating
func main() {
	arr := [...]int{1, 2, 1, 1, 2, 3, 4, 5, 6, 5, 6, 7, 8, 9, 8}
	countMap := make(map[int]int)

	//range returns index and element in case of arrays or slices
	for _, element := range arr {
		countMap[element] = countMap[element] + 1
	}

	//range returns key and values in case of maps
	for key, value := range countMap {
		fmt.Printf("element %d is repeated %d no.of times\n", key, value)
	}

	paragraph := "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?"

	arrayOfWords := strings.Fields(paragraph)

	mapOfWords := make(map[string]bool)

	for _, word := range arrayOfWords {
		if _, wordExist := mapOfWords[word]; !wordExist {
			mapOfWords[word] = true
		}
	}

	fmt.Println("total no of words (deduplicated) is ", len(mapOfWords))
}
