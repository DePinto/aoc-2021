package main

import (
	"reflect"
	"testing"
)

var testInput = []binaryFive{
	[5]int{0,0,1,0,0},
	[5]int{1,1,1,1,0},
	[5]int{1,0,1,1,0},
	[5]int{1,0,1,1,1},
	[5]int{1,0,1,0,1},
	[5]int{0,1,1,1,1},
	[5]int{0,0,1,1,1},
	[5]int{1,1,1,0,0},
	[5]int{1,0,0,0,0},
	[5]int{1,1,0,0,1},
	[5]int{0,0,0,1,0},
	[5]int{0,1,0,1,0},
}

var testInputFile = "day03_test.txt"

func TestParse(t *testing.T) {
	expected := testInput
	result := parse(testInputFile)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}