package main

import (
	"bufio"
	"fmt"
	"os"
)

var day = "05"

func file_to_lines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func max_int(haystack []int) (index, max int) {
	for i, v := range haystack {
		if max < v {
			max = v
			index = i
		}
	}
	return index, max
}

func sort_int(haystack []int) []int {
	var sorted []int

	for len(haystack) > 0 {
		i_max, max := max_int(haystack)
		sorted = append(sorted, max)
		// remove entry from haystack
		haystack[i_max] = haystack[len(haystack)-1]
		haystack = haystack[:len(haystack)-1]
	}
	return sorted
}

func binary_split(start_min, start_max int, steps []int) (min, max int) {
	min = start_min
	max = start_max
	for _, step := range steps {
		switch step {
		case -1: //up, left
			max = max - (max+1-min)/2
		case 1: //down, right
			min = min + (max-min)/2 + 1
		}
	}
	return min, max
}

func get_seat_position(code string) (row, col int) {
	var steps_row []int
	var steps_col []int

	for i, c := range code {
		if i < 7 {
			switch c {
			case 'F':
				steps_row = append(steps_row, -1)
			case 'B':
				steps_row = append(steps_row, 1)
			default:
				fmt.Println("illegal opcode: ", c)
			}
		} else {
			switch c {
			case 'L':
				steps_col = append(steps_col, -1)
			case 'R':
				steps_col = append(steps_col, 1)
			default:
				fmt.Println("illegal opcode: ", c)
			}
		}
	}

	min_row, max_row := binary_split(0, 127, steps_row)
	min_col, max_col := binary_split(0, 7, steps_col)

	if min_row != max_row || min_col != max_col {
		fmt.Println("Couldn't reduce code: ", code)
		os.Exit(1)
	}

	return min_row, min_col
}

func main() {
	fmt.Println("Adventofcode 2020 - day" + day)

	lines, err := file_to_lines("input/day" + day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// part01
	var seatIDs []int
	for _, line := range lines {
		row, col := get_seat_position(line)
		seatID := row*8 + col
		seatIDs = append(seatIDs, seatID)
	}
	_, maxSeat := max_int(seatIDs)
	fmt.Println("result part1: ", maxSeat)

	// part02
	myseat := -1
	seatIDs = sort_int(seatIDs)
	for i, seat := range seatIDs {
		if len(seatIDs)-i < 3 {
			break
		}

		if seatIDs[i+1] == seatIDs[i]-2 {
			myseat = seat - 1
		}
	}
	fmt.Println("result part2: ", myseat)
}
