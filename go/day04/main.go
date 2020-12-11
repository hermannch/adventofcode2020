package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var day = "04"

//type passport struct {
//	byr string
//	iyr string
//	eyr string
//	hgt string
//	hcl string
//	ecl string
//	pid string
//	cid string
//}

type passport map[string]string

func passport_has_required_fields(p passport) bool {
	_, byr_found := p["byr"]
	_, iyr_found := p["iyr"]
	_, eyr_found := p["eyr"]
	_, hgt_found := p["hgt"]
	_, hcl_found := p["hcl"]
	_, ecl_found := p["ecl"]
	_, pid_found := p["pid"]
	//_, found_cid := p["cid"]

	return byr_found && iyr_found && eyr_found && hgt_found && hcl_found && ecl_found && pid_found
}

func byr_valid(byr string) bool {
	val, err := strconv.Atoi(byr)
	if err == nil && val >= 1920 && val <= 2002 {
		return true
	}
	return false
}

func iyr_valid(iyr string) bool {
	val, err := strconv.Atoi(iyr)
	if err == nil && val >= 2010 && val <= 2020 {
		return true
	}
	return false
}

func eyr_valid(eyr string) bool {
	val, err := strconv.Atoi(eyr)
	if err == nil && val >= 2020 && val <= 2030 {
		return true
	}
	return false
}

func hgt_valid(hgt string) bool {
	if strings.HasSuffix(hgt, "cm") {
		hgt = strings.TrimSuffix(hgt, "cm")
		val, err := strconv.Atoi(hgt)
		if err == nil && val >= 150 && val <= 193 {
			return true
		}
	} else if strings.HasSuffix(hgt, "in") {
		hgt = strings.TrimSuffix(hgt, "in")
		val, err := strconv.Atoi(hgt)
		if err == nil && val >= 59 && val <= 76 {
			return true
		}
	}
	return false
}

func hcl_valid(hcl string) bool {
	if !strings.HasPrefix(hcl, "#") || len(hcl) != 7 {
		return false
	}
	hcl = strings.TrimPrefix(hcl, "#")
	for _, r := range strings.ToLower(hcl) {
		switch r {
		case 'a':
		case 'b':
		case 'c':
		case 'd':
		case 'e':
		case 'f':
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
		default:
			return false
		}
	}
	return true
}

func ecl_valid(ecl string) bool {
	switch ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	return true
}

func pid_valid(pid string) bool {
	_, err := strconv.Atoi(pid)
	if len(pid) != 9 || err != nil {
		return false
	}
	return true
}

func passport_has_valid_fields(p passport) bool {
	if !passport_has_required_fields(p) {
		return false
	}

	if !byr_valid(p["byr"]) {
		return false
	}

	if !iyr_valid(p["iyr"]) {
		return false
	}

	if !eyr_valid(p["eyr"]) {
		return false
	}

	if !hgt_valid(p["hgt"]) {
		return false
	}

	if !hcl_valid(p["hcl"]) {
		return false
	}

	if !ecl_valid(p["ecl"]) {
		return false
	}

	if !pid_valid(p["pid"]) {
		return false
	}

	return true
}

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

func lines_to_passports(lines []string) []passport {
	var passports []passport
	var p passport
	p = make(passport)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			for _, word := range strings.Fields(line) {
				attr := strings.Split(word, ":")
				p[attr[0]] = attr[1]
			}
		} else {
			passports = append(passports, p)
			p = make(passport)
		}
	}
	passports = append(passports, p) // don't forget the last entry :-)
	return passports
}

func main() {
	// load input data
	lines, err := file_to_lines("input/day" + day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := lines_to_passports(lines)

	// part01
	num_valid_passports1 := 0
	for _, p := range data {
		if passport_has_required_fields(p) {
			num_valid_passports1++
		}
	}

	// part02
	num_valid_passports2 := 0
	for _, p := range data {
		if passport_has_valid_fields(p) {
			num_valid_passports2++
		}
	}

	fmt.Println("Adventofcode 2020 - day", day, " part01: ", num_valid_passports1)
	fmt.Println("Adventofcode 2020 - day", day, " part01: ", num_valid_passports2)
}
