package main

import "fmt"

func main() {
	dict := make(map[string]string)

	fmt.Println(dict)

	dict["India"] = "Asia"
	dict["England"] = "Europe"

	fmt.Println(dict)
}
