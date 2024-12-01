package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1 := []int{}
	input2 := []int{}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		i1, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		input1 = append(input1, i1)
		input2 = append(input2, i2)
	}

	todo := len(input1)
	length := 0

	i1, x1, _ := findSmallest(input1)

	i2, x2, _ := findSmallest(input2)
	fmt.Println(i1, x1, i2, x2)

	for todo > 0 {
		todo--
		l, i1, i2, err := run(input1, input2)
		if err != nil {
			fmt.Println("Errror:", err)
			continue
		}

		input1 = i1
		input2 = i2
		length += l
		fmt.Println(l, length)
	}

	fmt.Println(length)

}

func run(input1 []int, input2 []int) (int, []int, []int, error) {
	i1, x1, err := findSmallest(input1)
	if err != nil {
		return 0, []int{}, []int{}, err
	}

	i2, x2, err := findSmallest(input2)
	if err != nil {
		return 0, []int{}, []int{}, err
	}

	res1 := input1[:i1]
	res1 = append(res1, input1[i1+1:]...)

	res2 := input2[:i2]
	res2 = append(res2, input2[i2+1:]...)

	len := 0
	if x1 < x2 {
		len = x2 - x1
	} else {
		len = x1 - x2
	}

	return len, res1, res2, nil
}

func findSmallest(input []int) (int, int, error) {
	var number int
	var index int

	if len(input) <= 0 {
		return 0, 0, fmt.Errorf("input empty")
	}

	for i, x := range input {
		if number == 0 || number >= x {
			number = x
			index = i
		}
	}

	return index, number, nil
}
