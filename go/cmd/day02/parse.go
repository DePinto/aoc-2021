package main

import (
	"strconv"
	"strings"
	"github.com/DePinto/aoc-2021/go/pkg/utils"
)

func parse(inputFilePath string) []move {
	xs, err := utils.NewStringSliceFromFileName(inputFilePath)
	if err != nil {
		panic(err)
	}

	ms := newMoveSlicefromStringSlice(xs)

	return ms
}

func newMoveSlicefromStringSlice(xs []string) []move {
	ms := []move{}

	for _, v := range xs {
		m := newMoveFromString(v)
		ms = append(ms, m)
	}

	return ms
}

func newMoveFromString(s string) move {
	xs := strings.Split(s, " ")
	direction := xs[0]
	distance, err := strconv.Atoi(xs[1])
	if err != nil {
		panic(err)
	}

	m := move{
		direction: direction,
		distance: distance,
	}

	return m
}