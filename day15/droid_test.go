package main

import (
	"aoc2019/common"
	"fmt"
	"testing"
)

func TestDroidWalkCompleteMap(t *testing.T) {
	d := newDroid()
	d.breadCrumbs = false
	d.haltOnOxySys = false
	err := d.cpu.Run()
	common.Assert(t, err == nil, "Failed to run program, error: %v", err)
	fmt.Println(d.getSpaceMap())
	op := pos{18, 18}
	common.Assert(t, d.oxySysPos == op, "oxySysPos = %v, want %v", d.oxySysPos, op)
	g, o := d.generateGrid()
	common.Assert(t, len(g) == 41, "Actual grid height = %d, want %d", len(g), 41)
	for i, r := range g {
		common.Assert(t, len(r) == 41, "Actual grid width = %d (row #%d), want %d", len(r), i, 41)
	}
	op = pos{39, 39}
	common.Assert(t, o == op, "Grid oxygen station = %v, want %v", o, op)
}

func TestPart1(t *testing.T) {
	d := part1()
	common.Assert(t, d != nil, "Failed to run program")
	m := d.getMinimumNumberOfMovements()
	common.Assert(t, m == 304, "Got %d movements, want %d", m, 304)
}

func TestPart2(t *testing.T) {
	l := part2()
	common.Assert(t, l == 310, "Part2 = %d, want %d", l, 310)
}
