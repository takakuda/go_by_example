package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
}
