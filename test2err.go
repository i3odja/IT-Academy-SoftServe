package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println("Incorrect Input")
	}
}