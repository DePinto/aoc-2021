package main

import (
	"strconv"
	"github.com/DePinto/aoc-2021/go/pkg/utils"
)

func parse(inputFile string) []binaryFive {
	xs, err := utils.NewStringSliceFromFileName(inputFile)
	if err != nil {
		panic(err)
	}

	bs := []binaryFive{}
	for _, s := range xs {
		b := newBinaryFiveFromString(s)
		bs = append(bs, b)
	}

	return bs
} 

func newBinaryFiveFromString(s string) binaryFive {
	b := binaryFive{}

	for i, v := range s {
		d, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}

		b[i] = d
	}

	return b
}