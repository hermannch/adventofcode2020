package main

import (
	"fmt"
	"testing"
)

var data = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func TestSumOfTwo(t *testing.T) {
	var sum = 2020
	var answer = 514579

	ind1, ind2, err := Find_sum_of_two(data, sum)
	if err != nil {
		t.Log(err)
		t.Fatal("fatal")
		t.Fail()
	}

	prod := int(data[ind1]) * int(data[ind2])
	if prod != answer {
		t.Log("ind1: " + fmt.Sprint(ind1))
		t.Log("ind2: " + fmt.Sprint(ind2))
		t.Fatal("")
		t.Fail()
	}
}

func TestSumOfThree(t *testing.T) {
	var sum = 2020
	var answer = 241861950

	ind1, ind2, ind3, err := Find_sum_of_three(data, sum)
	if err != nil {
		t.Log(err)
		t.Fatal("fatal")
		t.Fail()
	}

	prod := int(data[ind1]) * int(data[ind2]) * int(data[ind3])
	if prod != answer {
		t.Log("ind1: " + fmt.Sprint(ind1))
		t.Log("ind2: " + fmt.Sprint(ind2))
		t.Fatal("")
		t.Fail()
	}
}
