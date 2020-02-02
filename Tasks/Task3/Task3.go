package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func PythagoreanTriples(n int) [][3]int {
	result := make([][3]int, n)
	for i, y := 0, 0; i <= n; i++ {
		for j, a := 0, i; j <= n; j++ {
			for k, b := 0, j; k <= n; k++ {
				c := k
				if a != b && b != c && a != c {
					if a*a+b*b == c*c {
						if b >= a && b <= c {
							x := 0
							result[y][x] = a
							x++
							result[y][x] = b
							x++
							result[y][x] = c
							y++
						}
					}
				}
			}
		}
	}
	//cutting empty part of output
	for i, v := range result {
		if v[0] == 0 {
			result = result[0:i]
			return result
		}
	}
	return result
}

func fromFile(fileName string) {
	var input int

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
			continue
		}
		input = intText
		break
	}
	fmt.Println(input)
	result := PythagoreanTriples(input)
	fmt.Println(result)
}

func fromArgs(args []string) {
	var input int

	for _, v := range args {
		intInput, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Incorrect Input:", v)
			fmt.Println("Please, Enter correct number!")
			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				text := s.Text()
				if text == "" {
					break
				}
				intText, err := strconv.Atoi(text)
				if err != nil {
					fmt.Println("Incorrect Input", text)
					fmt.Println("enter:")
					continue
				}
				input = intText
				break
			}
		}
		if intInput == 0 {
			continue
		}
		input = intInput
	}
	result := PythagoreanTriples(input)
	fmt.Println(result)
}

func fromCommandLine() {
	var input int
	fmt.Println("Enter number [N]:")
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
		input = intText
		break
	}
	fmt.Println(input)
	result := PythagoreanTriples(input)
	fmt.Println(result)
}

var fileName = flag.String("file", "", "file path to read from")

func main() {
	fmt.Println("Pythagorean Triples")
	flag.Parse()
	switch {
	case *fileName != "":
		fromFile(*fileName)
	case len(os.Args[1:]) >= 1:
		fromArgs(os.Args[1:])
	default:
		fromCommandLine()
	}
}
