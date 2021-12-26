package main

import (
	"fmt"
	"github.com/DePinto/aoc-2021/go/pkg/utils"
)

const (
	inputFilePath = "/Users/ddepinto/go/src/github.com/DePinto/aoc-2021/input/day01.txt"
)

func main() {
	x := parse(inputFilePath)
	r1 := part1(x)
	r2 := part2(x)

	fmt.Printf("Part 1:\t%v\n", r1)
	fmt.Printf("Part 2:\t%v\n", r2)
}

func parse(inputFilePath string) []int {
	xs, err := utils.NewStringSliceFromFileName(inputFilePath)
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
	counter := 0

	for i, j := 0, 1; j < len(xi); i, j = i+1, j+1 {
		if xi[i] < xi[j] {
			counter++
		}
	}

	return counter
}

func part2(xi []int) int {
	counter := 0

	for i, j, k, l := 0, 1, 2, 3; l < len(xi); i, j, k, l = i+1, j+1, k+1, l+1 {
		sum1 := xi[i] + xi[j] + xi[k]
		sum2 := xi[j] + xi[k] + xi[l]
		if sum1 < sum2 {
			counter++
		}
	}

	return counter
}
