package main

import "math"

type intersection struct {
	x, y  int
	steps map[int]int
}

type wireGrid struct {
	data          [][]int
	sizeX         int
	centralPortX  int
	centralPortY  int
	wires         []*wire
	intersections map[int]intersection
}

func getRequiredSpace(wires []*wire) (x, y, centralPortX, centralPortY int) {
	x, y, centralPortX, centralPortY = 0, 0, 0, 0
	x1, x2, y1, y2 := 0, 0, 0, 0
	for _, w := range wires {
		if w.x1 < x1 {
			x1 = w.x1
		}
		if w.x2 > x2 {
			x2 = w.x2
		}
		if w.y1 < y1 {
			y1 = w.y1
		}
		if w.y2 > y2 {
			y2 = w.y2
		}
	}
	x = x2 - x1 + 1
	y = y2 - y1 + 1
	centralPortX = -x1
	centralPortY = -y1
	return
}

func (s *wireGrid) plotPath(id int, path []string) {
	cx, cy := s.centralPortX, s.centralPortY
	steps := 0
	addIntersection := func() {
		addr := cy*s.sizeX + cx
		is, ok := s.intersections[addr]
		if !ok {
			is = intersection{x: cx, y: cy, steps: make(map[int]int)}
		}
		is.steps[id] = steps
		s.intersections[addr] = is
	}
	draw := func(cxcy func(), l int) {
		for i := 0; i < l; i++ {
			cxcy()
			steps++
			v := s.data[cx][cy]
			if (v != 0) && (v != id) {
				addIntersection()
			}
			s.data[cx][cy] = id
		}
	}

	for _, p := range path {
		if d, l, e := getDirection(p); e == nil {
			switch d {
			case 'U':
				draw(func() { cy-- }, l)
			case 'R':
				draw(func() { cx++ }, l)
			case 'D':
				draw(func() { cy++ }, l)
			case 'L':
				draw(func() { cx-- }, l)
			}
		}
	}
}

func (s *wireGrid) plotWires() {
	for id, w := range s.wires {
		s.plotPath(id+1, w.path)
	}

	// Re-draw all except the last, to correctly detect where all wires intersects with other wires, and set correct step count for each wires intersection with other wires..
	for id := len(s.wires) - 2; id >= 0; id-- {
		s.plotPath(id+1, s.wires[id].path)
	}
}

func (s *wireGrid) getShortestManhattanDistance() int {
	abs := func(i int) int {
		if i >= 0 {
			return i
		}
		return -i
	}
	d := math.MaxInt32
	for _, is := range s.intersections {
		dc := abs(s.centralPortX-is.x) + abs(s.centralPortY-is.y)
		if dc < d {
			d = dc
		}
	}
	return d
}

func (s *wireGrid) getShortestStepPath() int {
	d := math.MaxInt32
	for _, is := range s.intersections {
		sc := 0
		for _, s := range is.steps {
			sc += s
		}
		if sc < d {
			d = sc
		}
	}
	return d
}

func newWireGrid(wires []*wire) *wireGrid {
	sizeX, sizeY, centralPortX, centralPortY := getRequiredSpace(wires)
	s := &wireGrid{
		data:          make([][]int, sizeX),
		sizeX:         sizeX,
		centralPortX:  centralPortX,
		centralPortY:  centralPortY,
		wires:         wires,
		intersections: make(map[int]intersection),
	}

	// Allocate the rest of the wire grid..
	for x := 0; x < sizeX; x++ {
		s.data[x] = make([]int, sizeY)
	}

	// Plot all wires onto the grid
	s.plotWires()

	return s
}
