package main

import (
	"fmt"
)

const (
	inputFile = "day02.txt"
)

type move struct {
	direction string
	distance  int
}

func main() {
	moves := parse(inputFile)
	r1 := part1(moves)
	r2 := part2(moves)

	fmt.Printf("Part 1:\t%v\n", r1)
	fmt.Printf("Part 2:\t%v\n", r2)
}
