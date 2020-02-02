package main

import (
	"fmt"
	"os"
	"strconv"
)

func insertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
}

func main() {
	args := os.Args[1:]
	a := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		fmt.Println(i)
		a[i], _ = strconv.Atoi(args[i])
	}
	//a := []int{5,4,19,6,10,1,102,45,9,235,7,2,3}
	fmt.Printf("Unsorted: \t%v", a)
	insertionSort(a)
	fmt.Printf("\nSorted: \t%v", a)
}
