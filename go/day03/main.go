package main

import (
	"bufio"
	"fmt"
	"os"
)

var day = "03"

type Slope struct {
	x int
	y int
}

func count_trees_hit_with_slope(data []string, slope Slope) int {
	var row int = 0
	var num_trees int = 0

	for line := slope.y; line < len(data); line += slope.y {
		row += slope.x

		// since the patterns repeat on the right, just reset the row index in
		// case of overflow
		for row >= len(data[line]) {
			row = row - len(data[line])
		}

		if string(data[line][row]) == "#" {
			num_trees += 1
		}
	}

	return num_trees
}

func count_trees_hit_with_multiple_slopes(data []string, slopes []Slope) int {
	var prod int = 1
	for _, s := range slopes {
		res := count_trees_hit_with_slope(data, s)
		prod = prod * res
	}
	return prod
}

func main() {
	fmt.Println("Adventofcode 2020 - day" + day)

	// load input data
	file, err := os.Open("input/day" + day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var data = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	// part01
	s1 := Slope{3, 1}
	fmt.Println("result part1: ", count_trees_hit_with_slope(data, s1))

	// part02
	s2 := []Slope{
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}

	fmt.Println("result part2: ", count_trees_hit_with_multiple_slopes(data, s2))
}
