package main

import (
	//"fmt"
	"github.com/DePinto/aoc-2021/go/pkg/utils"
)

const (
	inputFilePath = "/Users/ddepinto/go/src/github.com/DePinto/aoc-2021/input/day01.txt"
	testInputFilePath = "/Users/ddepinto/go/src/github.com/DePinto/aoc-2021/input/day01_test.txt"
)

func main() {

}

func parse(inputFilePath string) []int {
	xs, err := utils.NewStringSliceFromFileName(testInputFilePath)
	if err != nil {
		panic(err)
	}

	xi, err := utils.NewIntSliceFromStringSlice(xs)
	if err != nil {
		panic(err)
	}

	return xi
}


func part1(xi []int) int {
	return -1
}