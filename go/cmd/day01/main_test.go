package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 7
	result := part1([]string{})

	if result != expected {
		t.Errorf("\tgot:\t%v\n\t\twanted:\t%v\n", result, expected)
	}
}