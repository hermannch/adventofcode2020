package main

import (
	"testing"
)

var data_raw_1 = []string{
	"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
	"byr:1937 iyr:2017 cid:147 hgt:183cm",
	"",
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
	"hcl:#cfa07d byr:1929",
	"",
	"hcl:#ae17e1 iyr:2013",
	"eyr:2024",
	"ecl:brn pid:760753108 byr:1931",
	"hgt:179cm",
	"",
	"hcl:#cfa07d eyr:2025 pid:166559648",
	"iyr:2011 ecl:brn hgt:59in",
}

var data_raw2_invalid = []string{
	"eyr:1972 cid:100",
	"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
	"",
	"iyr:2019",
	"hcl:#602927 eyr:1967 hgt:170cm",
	"ecl:grn pid:012533040 byr:1946",
	"",
	"hcl:dab227 iyr:2012",
	"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
	"",
	"hgt:59cm ecl:zzz",
	"eyr:2038 hcl:74454a iyr:2023",
	"pid:3556412378 byr:2007",
}

var data_raw2_valid = []string{
	"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
	"hcl:#623a2f",
	"",
	"eyr:2029 ecl:blu cid:129 byr:1989",
	"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
	"",
	"hcl:#888785",
	"hgt:164cm byr:2001 iyr:2015 cid:88",
	"pid:545766238 ecl:hzl",
	"eyr:2022",
	"",
	"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
}

func Test_part1(t *testing.T) {
	data := lines_to_passports(data_raw_1)

	if !passport_has_required_fields(data[0]) {
		t.Fatal(0, ": expected valid passport")
	}
	if passport_has_required_fields(data[1]) {
		t.Fatal(1, ": expected invalid passport")
	}
	if !passport_has_required_fields(data[2]) {
		t.Fatal(2, ": expected valid passport")
	}
	if passport_has_required_fields(data[3]) {
		t.Fatal(3, ": expected invalid passport")
	}
}

func Test_part2(t *testing.T) {

	if !byr_valid("2002") {
		t.Fatal("Exptected 2002 to be valid")
	}
	if byr_valid("2003") {
		t.Fatal("Exptected 2003 to be invalid")
	}

	if !hgt_valid("60in") {
		t.Fatal("Exptected 60in to be valid")
	}
	if !hgt_valid("190cm") {
		t.Fatal("Exptected 190cm to be valid")
	}
	if hgt_valid("190in") {
		t.Fatal("Exptected 190in to be invalid")
	}
	if hgt_valid("190") {
		t.Fatal("Exptected 190 to be invalid")
	}

	if !hcl_valid("#123abc") {
		t.Fatal("Exptected #123abc to be valid")
	}
	if hcl_valid("#123abz") {
		t.Fatal("Exptected #123abz to be invalid")
	}
	if hcl_valid("123abc") {
		t.Fatal("Exptected 123abc to be invalid")
	}

	if !ecl_valid("brn") {
		t.Fatal("Exptected brn to be valid")
	}
	if ecl_valid("wat") {
		t.Fatal("Exptected wat to be invalid")
	}

	if !pid_valid("000000001") {
		t.Fatal("Exptected pid:000000001 to be valid")
	}
	if pid_valid("0123456789") {
		t.Fatal("Exptected pid:0123456789 to be invalid")
	}

	for i, p := range lines_to_passports(data_raw2_invalid) {
		if passport_has_valid_fields(p) {
			t.Fatal(i, ": expected invalid passport")
		}
	}

	for i, p := range lines_to_passports(data_raw2_valid) {
		if !passport_has_valid_fields(p) {
			t.Fatal(i, ": expected valid passport")
		}
	}
}
