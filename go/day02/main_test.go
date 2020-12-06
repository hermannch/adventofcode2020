package main

import (
	"fmt"
	"testing"
)

var data = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

func TestPolicy1(t *testing.T) {
	var e DBentry

	e, _ = convert_string_to_DBentry(data[0])
	if !pw_matches_policy1(e) {
		t.Fatal("0")
	}

	e, _ = convert_string_to_DBentry(data[1])
	if pw_matches_policy1(e) {
		t.Fatal("1")
	}

	e, _ = convert_string_to_DBentry(data[2])
	if !pw_matches_policy1(e) {
		t.Fatal("2")
	}
}

func TestPolicy2(t *testing.T) {
	var e DBentry

	e, _ = convert_string_to_DBentry(data[0])
	if !pw_matches_policy2(e) {
		t.Fatal("0")
	}

	e, _ = convert_string_to_DBentry(data[1])
	if pw_matches_policy2(e) {
		t.Fatal("1")
	}

	e, _ = convert_string_to_DBentry(data[2])
	if pw_matches_policy2(e) {
		t.Fatal("2")
	}
}

func test_convert_string_to_DBentry(t *testing.T, input string, answer DBentry) {
	res, err := convert_string_to_DBentry(input)
	if err != nil {
		t.Log(err)
		t.Fatal("fatal")
		t.Fail()
	}

	if answer != res {
		t.Log("answer: " + fmt.Sprint(answer))
		t.Log("res: " + fmt.Sprint(res))
		t.Fatal("")
		t.Fail()
	}
}
func TestConvert_string_to_DBentry(t *testing.T) {

	var input string
	var answer DBentry

	input = "1-3 a: abcde"
	answer.policy.min = 1
	answer.policy.max = 3
	answer.policy.char = "a"[0]
	answer.password = "abcde"
	test_convert_string_to_DBentry(t, input, answer)

	input = "1-3 b: cdefg"
	answer.policy.min = 1
	answer.policy.max = 3
	answer.policy.char = "b"[0]
	answer.password = "cdefg"
	test_convert_string_to_DBentry(t, input, answer)

	input = "1-3 a: abcde"
	answer.policy.min = 1
	answer.policy.max = 3
	answer.policy.char = "a"[0]
	answer.password = "abcde"
	test_convert_string_to_DBentry(t, input, answer)
}
