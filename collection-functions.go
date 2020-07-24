package main

import (
	"fmt"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
}
