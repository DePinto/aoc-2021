package main

import (
	"reflect"
	"testing"
)

var day01TestInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
var day01TestInputFile = "day01_test.txt"
var day01TestResultPartOne = 7
var day01TestResultPartTwo = 5

func TestParse(t *testing.T) {
	expected := day01TestInput
	result := parse(day01TestInputFile)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestPart1(t *testing.T) {
	expected := day01TestResultPartOne
	result := part1(day01TestInput)

	if result != expected{
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := day01TestResultPartTwo
	result := part2(day01TestInput)

	if result != expected{
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}
