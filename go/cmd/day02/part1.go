package main

import ()

type position struct {
	horizontial int
	depth       int
}

func part1(ms []move) int {
	p := position{0, 0}

	for _, v := range ms {
		p.move(v)
	}

	return p.horizontial * p.depth
}

func (p *position) move(m move) *position {
	switch m.direction {
	case "forward":
		p.horizontial = p.horizontial + m.distance
	case "down":
		p.depth = p.depth + m.distance
	case "up":
		p.depth = p.depth - m.distance
	}

	return p
}
