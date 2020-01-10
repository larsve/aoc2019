package main

import (
	"fmt"
	"testing"

	"aoc2019/common"
)

func TestMove(t *testing.T) {
	r := &robot{pos: pos{0, 0}, dir: 0, out: 0, grid: make(map[pos]int)}
	tt := []struct {
		sdir, sx, sy int // Starting values
		mdir         int // Move direction (left / right)
		edir, ex, ey int // Expected values
	}{
		{sdir: 0, sx: 0, sy: 0, mdir: 1, edir: 1, ex: 1, ey: 0},   // Up -> Right
		{sdir: 1, sx: 1, sy: 0, mdir: 1, edir: 2, ex: 1, ey: 1},   // Right -> Down
		{sdir: 2, sx: 1, sy: 1, mdir: 1, edir: 3, ex: 0, ey: 1},   // Down -> Left
		{sdir: 3, sx: 0, sy: 1, mdir: 1, edir: 0, ex: 0, ey: 0},   // Left -> Up
		{sdir: 0, sx: 0, sy: 0, mdir: 0, edir: 3, ex: -1, ey: 0},  // Up -> Left
		{sdir: 3, sx: -1, sy: 0, mdir: 0, edir: 2, ex: -1, ey: 1}, // Left -> Down
		{sdir: 2, sx: -1, sy: 1, mdir: 0, edir: 1, ex: 0, ey: 1},  // Down -> Right
		{sdir: 1, sx: 0, sy: 1, mdir: 0, edir: 0, ex: 0, ey: 0},   // Right -> Up
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test#%v", i), func(t *testing.T) {
			r.dir = tc.sdir
			r.pos.x = tc.sx
			r.pos.y = tc.sy
			r.move(tc.mdir)
			common.Assert(t, r.dir == tc.edir, "Not the expected direction, expected %v but got %v", tc.edir, r.dir)
			common.Assert(t, r.pos.x == tc.ex, "Not the expected X position, expected %v but got %v", tc.ex, r.pos.x)
			common.Assert(t, r.pos.y == tc.ey, "Not the expected Y position, expected %v but got %v", tc.ey, r.pos.y)
		})
	}
}

func TestPart1(t *testing.T) {
	r := part1()
	common.Assert(t, r != nil, "Unexpected error, result is nil")
	l := len(r.grid)
	common.Assert(t, l == 2018, "Unexpected length, expected 2018 but got %v", l)
	common.Assert(t, r.topLeft == pos{-14, -33}, "Unexpected top/left, expected -14/-33 but got %v", r.topLeft)
	common.Assert(t, r.bottomRight == pos{41, 67}, "Unexpected bottom/right, expected 41/67 but got %v", r.bottomRight)
}

func TestPart2(t *testing.T) {
	r := part2()
	common.Assert(t, r != nil, "Unexpected error, result is nil")
	l := len(r.grid)
	common.Assert(t, l == 249, "Unexpected length, expected 2018 but got %v", l)
	common.Assert(t, r.topLeft == pos{0, 0}, "Unexpected top/left, expected -14/-33 but got %v", r.topLeft)
	common.Assert(t, r.bottomRight == pos{42, 5}, "Unexpected bottom/right, expected 41/67 but got %v", r.bottomRight)
}

func BenchmarkPaint(b *testing.B) {
	r := newHullPainterRobot()
	r.grid[pos{0, 0}] = 1 // Start on a white panel
	for n := 0; n < b.N; n++ {
		if err := r.run(); err != nil {
			fmt.Println("Paint panels failed with error:", err)
		}
	}

}

func BenchmarkGetPanels(b *testing.B) {
	r := newHullPainterRobot()
	r.grid[pos{0, 0}] = 1 // Start on a white panel
	if err := r.run(); err != nil {
		fmt.Println("Paint panels failed with error:", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.getPanels()
	}
}
