package main

import (
	"reflect"
	"testing"
)

var testInput = []move{
	{
		direction: "forward",
		distance:  5,
	},
	{
		direction: "down",
		distance:  5,
	},
	{
		direction: "forward",
		distance:  8,
	},
	{
		direction: "up",
		distance:  3,
	},
	{
		direction: "down",
		distance:  8,
	},
	{
		direction: "forward",
		distance:  2,
	},
}

var testInputFile = "day02_test.txt"
var testResultPartOne = 150
var testResultPartTwo = 900

func TestParse(t *testing.T) {
	expected := testInput
	result := parse(testInputFile)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestPart1(t *testing.T) {
	expected := testResultPartOne
	result := part1(testInput)

	if result != expected {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := testResultPartTwo
	result := part2(testInput)

	if result != expected {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}
