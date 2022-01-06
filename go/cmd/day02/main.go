package main

import(
	"fmt"
)

const (
	inputFilePath = "C:\\Users\\danie\\go\\src\\github.com\\DePinto\\aoc-2021\\input\\day02.txt"
)

type move struct{
	direction string
	distance int
}

func main() {
	x := parse(inputFilePath)
	r1 := part1(x)
	r2 := part2(x)

	fmt.Printf("Part 1:\t%v\n", r1)
	fmt.Printf("Part 2:\t%v\n", r2)
}

