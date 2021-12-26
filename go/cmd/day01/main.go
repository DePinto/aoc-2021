package main

import (
	//"fmt"
	"github.com/DePinto/aoc-2021/go/pkg/utils"
)

const inputFilePath = "/Users/ddepinto/go/src/github.com/DePinto/aoc-2021/input/day01_test.txt"

func main() {

}

func Parse(inputFilePath string) []int {
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


func Part1(xi []int) int {
	return -1
}