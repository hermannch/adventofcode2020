package main

import (
	"testing"
)

func Test_binary_split(t *testing.T) {
	var min, max int

	min, max = binary_split(0, 15, []int{-1})
	if min != 0 {
		t.Fatal("FAIL")
	}
	if max != 7 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{1})
	if min != 8 {
		t.Fatal("FAIL (got ", min, ")")
	}
	if max != 15 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{-1, -1})
	if min != 0 {
		t.Fatal("FAIL")
	}
	if max != 3 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{-1, 1})
	if min != 4 {
		t.Fatal("FAIL")
	}
	if max != 7 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{1, -1})
	if min != 8 {
		t.Fatal("FAIL")
	}
	if max != 11 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{1, 1})
	if min != 12 {
		t.Fatal("FAIL")
	}
	if max != 15 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{-1, -1, -1})
	if min != 0 {
		t.Fatal("FAIL")
	}
	if max != 1 {
		t.Fatal("FAIL (got ", max, ")")
	}

	min, max = binary_split(0, 15, []int{-1, -1, 1})
	if min != 2 {
		t.Fatal("FAIL")
	}
	if max != 3 {
		t.Fatal("FAIL (got ", max, ")")
	}
}

func Test_part1(t *testing.T) {
	row, col := get_seat_position("FBFBBFFRLR")
	if row != 44 {
		t.Fatal("expected 44, got ", row)
	}
	if col != 5 {
		t.Fatal("expected 5, got ", row)
	}
}
