package main

import (
	"testing"
)

var data = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}
var answer1 int = 7

var s2 []Slope = []Slope{
	Slope{1, 1},
	Slope{3, 1},
	Slope{5, 1},
	Slope{7, 1},
	Slope{1, 2},
}
var a2 []int = []int{2, 7, 3, 4, 2}
var answer2 int = 336

func Test_part1(t *testing.T) {

	res := count_trees_hit_with_slope(data, Slope{3, 1})
	if res != answer1 {
		t.Log("got ", res)
		t.Fatal("expected ", answer1)
	}
}

func Test_part2(t *testing.T) {

	var s Slope
	var a int
	var res int

	s = Slope{1, 1}
	a = 2
	res = count_trees_hit_with_slope(data, s)
	if res != a {
		t.Fatal("expected ", a, " but got", res)
	}

	s = Slope{3, 1}
	a = 7
	res = count_trees_hit_with_slope(data, s)
	if res != a {
		t.Fatal("expected ", a, " but got", res)
	}

	s = Slope{5, 1}
	a = 3
	res = count_trees_hit_with_slope(data, s)
	if res != a {
		t.Fatal("expected ", a, " but got", res)
	}

	s = Slope{7, 1}
	a = 4
	res = count_trees_hit_with_slope(data, s)
	if res != a {
		t.Fatal("expected ", a, " but got", res)
	}

	s = Slope{1, 2}
	a = 2
	res = count_trees_hit_with_slope(data, s)
	if res != a {
		t.Fatal("expected ", a, " but got", res)
	}

	res = count_trees_hit_with_multiple_slopes(data, s2)
	if res != answer2 {
		t.Fatal("got ", res)
		t.Fatal("expected ", answer2)
	}
}
