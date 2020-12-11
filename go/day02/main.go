package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var day = "02"

type Policy struct {
	char byte
	min  int
	max  int
}

type DBentry struct {
	policy   Policy
	password string
}

// extract policy and pw from raw input
func convert_string_to_DBentry(input string) (DBentry, error) {
	var e DBentry
	s := strings.Split(input, ":")
	e.password = strings.TrimSpace(s[1])

	s = strings.Split(s[0], " ")
	e.policy.char = s[1][0]

	s = strings.Split(s[0], "-")
	e.policy.min, _ = strconv.Atoi(s[0])
	e.policy.max, _ = strconv.Atoi(s[1])

	return e, nil
}

func pw_matches_policy1(input DBentry) bool {
	var num = 0
	for _, c := range input.password {
		if byte(c) == input.policy.char {
			num++
		}
	}
	if num <= input.policy.max && num >= input.policy.min {
		return true
	}
	return false
}

func pw_matches_policy2(input DBentry) bool {
	first := input.password[input.policy.min-1]
	second := input.password[input.policy.max-1]
	char := input.policy.char

	if (first == char || second == char) && first != second {
		return true
	}
	return false
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

	var data = []DBentry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e, err := convert_string_to_DBentry(scanner.Text())
		if err == nil {
			data = append(data, e)
		}
	}

	var valid1 int
	var invalid1 int
	for _, val := range data {
		if pw_matches_policy1(val) {
			valid1++
		} else {
			invalid1++
		}
	}
	fmt.Println("resut part1: %i", valid1)

	var valid2 int
	for _, val := range data {
		if pw_matches_policy2(val) {
			valid2++
		}
	}
	fmt.Println("resut part2: %i", valid2)
}
