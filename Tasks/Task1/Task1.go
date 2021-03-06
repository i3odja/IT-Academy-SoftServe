package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Multiplicity3not5(input []int) int {
	var count int
	for _, v := range input {
		if v%3 == 0 && v%5 != 0 {
			count++
		}
	}
	return count
}

func fromFile(fileName string) {
	var input []int

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Incorrect Input")
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for s.Scan() {
		text := s.Text()
		intText, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Incorrect Input", text)
		}
		input = append(input, intText)
	}
	fmt.Println(input)
	result := Multiplicity3not5(input)
	fmt.Println(result)
}

func fromArgs(args []string) {

	var input []int

	for _, v := range args {
		intInput, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Incorrect Input:", v)
		}
		if intInput == 0 {
			continue
		}
		input = append(input, intInput)
	}
	fmt.Println(input)
	result := Multiplicity3not5(input)
	fmt.Println(result)
}

func fromCommandLine() {
	var input []int
	fmt.Println("Enter numbers:")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		if text == "" {
			break
		}
		intText, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Incorrect Input", text)
			continue
		}
		input = append(input, intText)
	}
	fmt.Println(input)
	result := Multiplicity3not5(input)
	fmt.Println(result)
}

var fileName = flag.String("file", "", "file path to read from")

func main() {
	flag.Parse()
	switch {
	case *fileName != "":
		fromFile(*fileName)
	case len(os.Args[1:]) > 1:
		fromArgs(os.Args[1:])
	default:
		fromCommandLine()
	}
}
