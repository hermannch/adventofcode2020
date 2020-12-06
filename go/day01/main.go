package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var day = "01"

func Find_sum_of_two(data []int, needle_sum int) (int, int, error) {

	ind1 := 0
	ind2 := 0
	found := false

	// naive bubble-sort like walk over all elements
	for ind1 = 0; ind1 < len(data) && !found; ind1++ {
		for ind2 = ind1 + 1; ind2 < len(data) && !found; ind2++ {
			//fmt.Println(ind1, ind2)
			if data[ind1]+data[ind2] == needle_sum {
				found = true
			}
		}
	}
	// fix off-by-one due to post command
	ind1 -= 1
	ind2 -= 1

	if found {
		return ind1, ind2, nil
	} else {
		return 0, 0, errors.New("needle not found")
	}
}

func Find_sum_of_three(data []int, needle_sum int) (int, int, int, error) {

	ind1 := 0
	ind2 := 0
	ind3 := 0
	found := false

	// naive bubble-sort like walk over all elements
	for ind1 = 0; ind1 < len(data) && !found; ind1++ {
		for ind2 = ind1 + 1; ind2 < len(data) && !found; ind2++ {
			for ind3 = ind2 + 1; ind3 < len(data) && !found; ind3++ {
				if data[ind1]+data[ind2]+data[ind3] == needle_sum {
					found = true
				}
			}
		}
	}
	// fix off-by-one due to post command
	ind1 -= 1
	ind2 -= 1
	ind3 -= 1

	if found {
		return ind1, ind2, ind3, nil
	} else {
		return 0, 0, 0, errors.New("needle not found")
	}
}

func main() {
	fmt.Println("Adventofcode 2020 - day" + day)
	sum_target := 2020

	//load values from file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var data = []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err == nil {
			data = append(data, num)
		}
	}

	// part1
	ind1, ind2, err := Find_sum_of_two(data, sum_target)
	if err != nil {
		panic(err)
	}
	prod2 := int(data[ind1]) * int(data[ind2])
	fmt.Println("result part1: ", prod2)

	// part2
	ind1, ind2, ind3, err := Find_sum_of_three(data, sum_target)
	if err != nil {
		panic(err)
	}
	prod3 := int(data[ind1]) * int(data[ind2]) * int(data[ind3])
	fmt.Println("result part2: ", prod3)
}
