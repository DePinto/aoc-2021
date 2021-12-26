package main

import (
	"reflect"
	"testing"
)

const (
	day01_test_result = 7
	testInputFilePath = "/Users/ddepinto/go/src/github.com/DePinto/aoc-2021/input/day01_test.txt"
)

func TestParse(t *testing.T) {
	expected := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	result := Parse(testInputFilePath)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestPart1(t *testing.T) {
	expected := day01_test_result
	result := Part1([]int{})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}