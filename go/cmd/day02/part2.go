package main

import ()

type positionWithAim struct {
	position
	aim int
}

func part2(ms []move) int {
	p := positionWithAim{
		position: position{
			horizontial: 0,
			depth:       0,
		},
		aim: 0,
	}

	for _, v := range ms {
		p.move(v)
	}

	return p.horizontial * p.depth
}

func (p *positionWithAim) move(m move) *positionWithAim {
	switch m.direction {
	case "forward":
		p.horizontial = p.horizontial + m.distance
		p.depth = p.depth + (p.aim * m.distance)
	case "down":
		p.aim = p.aim + m.distance
	case "up":
		p.aim = p.aim - m.distance
	}

	return p
}
