package main

import (
	"reflect"
	"testing"
)

var day01TestInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
var day01TestInputFilePath = "/Users/ddepinto/go/src/github.com/DePinto/aoc-2021/input/day01_test.txt"
var day01TestResult = 7

func TestParse(t *testing.T) {
	expected := day01TestInput
	result := parse(day01TestInputFilePath)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}

func TestPart1(t *testing.T) {
	expected := day01TestResult
	result := part1([]int{})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}
