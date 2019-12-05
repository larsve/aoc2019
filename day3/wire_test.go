package main

import (
	"testing"

	"aoc2019/common"
)

func TestGetDirection(t *testing.T) {
	test := func(direction string, edir byte, elen int) {
		d, l, e := getDirection(direction)
		common.Assert(t, e == nil, "Direction %v caused an error %v", direction, e)
		common.Assert(t, d == edir, "Direction %v did not result in expected dir %v, got %v", direction, edir, d)
		common.Assert(t, d == edir, "Direction %v did not result in expected len %v, got %v", direction, elen, l)
	}
	test("U83", 'U', 83)
	test("R75", 'R', 75)
	test("D30", 'D', 30)
	test("L72", 'L', 72)
}

func TestWireGetRequiredBoardSize(t *testing.T) {
	w:= newWire("U5,R5,D10,L10")
	common.Assert(t, w.x1 == -5, "X1 not the expected value, got %v", w.x1)
	common.Assert(t, w.x2 == +5, "X2 not the expected value, got %v", w.x2)
	common.Assert(t, w.y1 == -5, "Y1 not the expected value, got %v", w.y1)
	common.Assert(t, w.y2 == +5, "Y2 not the expected value, got %v", w.y2)
}
