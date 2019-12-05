package main

import (
	"strings"
	"testing"

	"aoc2019/common"
)

func TestGetRequiredWireSpace(t *testing.T) {
	wires := []*wire{}
	test := func(directions string, ex, ey, ecx, ecy int) {
		wires = append(wires, newWire(directions))
		x, y, cx, cy := getRequiredSpace(wires)
		common.Assert(t, x == ex, "Not expected X, expected %v, but got %v", ex, x)
		common.Assert(t, y == ey, "Not expected Y, expected %v, but got %v", ey, y)
		common.Assert(t, cx == ecx, "Not expected CX, expected %v, but got %v", ecx, cx)
		common.Assert(t, cy == ecy, "Not expected CY, expected %v, but got %v", ecy, cy)
	}
	test("R5,D10", 6, 11, 0, 0)
	test("R98,U47,R26,D63", 125, 64, 0, 47)
	test("L42", 167, 64, 42, 47)
	test(wire1, 11280, 8922, 1450, 5446)
	test(wire2, 16854, 14767, 7024, 5446)
}

func TestPlotPath(t *testing.T) {
	s := newWireGrid([]*wire{newWire("U5,R5,D10,L10")})
	test := func(id int, path string) int {
		s.plotPath(id, strings.Split(path, ","))
		return len(s.intersections)
	}

	cnt := test(2, "R4,D2,L2,U4")
	common.Assert(t, cnt == 0, "Got intersections for self crossing wire:%v", s)
	cnt = test(3, "D4,R3,U4")
	common.Assert(t, cnt == 2, "Expected two intersection, but got %v (%v)", cnt, s.intersections)
	cnt = test(4, "L5,D4,R6,U8,R2,D4")
	common.Assert(t, cnt == 5, "Expected five intersections, but got %v (%v)", cnt, s.intersections)
}

func testWires1() []*wire {
	return []*wire{
		newWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
		newWire("U62,R66,U55,R34,D71,R55,D58,R83"),
	}
}

func testWires2() []*wire {
	return []*wire{
		newWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
		newWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
	}
}

func TestTestData1(t *testing.T) {
	g := newWireGrid(testWires1())
	d := g.getShortestManhattanDistance()
	common.Assert(t, d == 159, "Not the expected distance, got %d", d)
	d = g.getShortestStepPath()
	common.Assert(t, d == 610, "Not the expected step count, got %d", d)
}
func TestTestData2(t *testing.T) {
	g := newWireGrid(testWires2())
	d := g.getShortestManhattanDistance()
	common.Assert(t, d == 135, "Not the expected distance, got %d", d)
	d = g.getShortestStepPath()
	common.Assert(t, d == 410, "Not the expected step count, got %d", d)
}

func TestPuzzle(t *testing.T) {
	g := newWireGrid([]*wire{newWire(wire1), newWire(wire2)})
	d := g.getShortestManhattanDistance()
	common.Assert(t, d == 232, "Not the expected distance, got %d", d)
	d = g.getShortestStepPath()
	common.Assert(t, d == 6084, "Not the expected step count, got %d", d)
}

func BenchmarkDay3Solutions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		g := newWireGrid([]*wire{newWire(wire1), newWire(wire2)})
		_ = g.getShortestManhattanDistance()
		_ = g.getShortestStepPath()
	}
}
