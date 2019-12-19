package main

import (
	"bytes"
	"fmt"
	"testing"

	"aoc2019/common"
)

var (
	// 5 X 5, Best @3,4 = 8
	example1 = ".#..#.....#####....#...##"

	// 10 X 10, Best @5,8 = 33
	example2 = "......#.#.#..#.#......#######..#.#.###...#..#.......#....#.##..#....#..##.#..#####...#..#..#....####"

	// 10 X 10, Best @1,2 = 35
	example3 = "#.#...#.#..###....#..#....#...##.#.#.#.#....#.#.#..##..###.#..#...##....##....##......#....####.###."

	// 10 X 10, Best @6,3 = 41
	example4 = ".#..#..#######.###.#....###.#...###.##.###.##.#.#.....###..#..#.#..#.##..#.#.###.##...##.#.....#.#.."

	// 20 X 20, Best @11,13 = 210
	example5 = ".#..##.###...#########.############..##..#.######.########.#.###.#######.####.#.#####.##.#.##.###.##..#####..#.##############################.####....###.#.#.####.######################.##.###..####....######..##.###########.##.####...##..#.#####..#.######.#####...#.##########...#.##########.#######.####.#.###.###.#.##....##.##.###..#####.#.#.###########.####.#.#.#####.####.######.##.####.##.#..##"
)

func check(t *testing.T, width, height int, data string, acount, bestX, bestY, bestCnt int) {
	m, e := newMap(width, height, bytes.NewBufferString(data))
	common.Assert(t, e == nil, "Failed to create map: %v", e)
	common.Assert(t, m != nil, "Returned nil")
	checkMap(t, m, acount, bestX, bestY, bestCnt)
}

func checkMap(t *testing.T, amap *asteroidMap, acount, bestX, bestY, bestCnt int) {
	common.Assert(t, len(amap.asteroids) == acount, "Not the expected number of asteroids, expected %v but got %v", acount, len(amap.asteroids))
	x, y, c := amap.getBestPlacement()
	common.Assert(t, x == bestX, "Not the expected X coordinate, expected %v but got %v", bestX, x)
	common.Assert(t, y == bestY, "Not the expected Y coordinate, expected %v but got %v", bestY, y)
	common.Assert(t, c == bestCnt, "Not the expected count, expected %v but got %v", bestCnt, c)
}

func TestGetVector(t *testing.T) {
	am := &asteroidMap{width: 21, height: 21}
	validate := func(x1, y1, x2, y2, xv, yv int) {
		av := am.getVector(x1, y1, x2, y2)
		ev := yv*am.width + xv
		common.Assert(t, av == ev, "Not the expected vector for %v, %v - %v, %v, expected %v but got %v", x1, y1, x2, y2, ev, av)
	}
	validate(10, 10, 10, 10, 0, 0)
	validate(10, 10, 20, 10, 1, 0)
	validate(10, 10, 10, 20, 0, 1)
	validate(10, 10, 0, 10, -1, 0)
	validate(10, 10, 10, 0, 0, -1)

	validate(10, 10, 20, 0, 1, -1)
	validate(10, 10, 20, 20, 1, 1)
	validate(10, 10, 0, 20, -1, 1)
	validate(10, 10, 0, 0, -1, -1)

	validate(10, 10, 13, 19, 1, 3)
	validate(10, 10, 13, 20, 3, 10)
	validate(10, 10, 18, 3, 8, -7)
	validate(10, 10, 18, 2, 1, -1)
	validate(10, 10, 18, 1, 8, -9)

	validate(0, 1, 3, 3, 3, 2)
	validate(0, 1, 9, 7, 3, 2)

	validate(0, 1, 4, 4, 4, 3)
	validate(0, 1, 8, 7, 4, 3)
}

func TestAngle(t *testing.T) {
	tt := []struct {
		cx, cy int     // Center point
		tx, ty int     // Target point
		deg    float64 // Expected degrees
	}{
		{cx: 0, cy: 0, tx: 0, ty: -5, deg: 0},
		{cx: 0, cy: 0, tx: 5, ty: -5, deg: 45},
		{cx: 0, cy: 0, tx: 5, ty: 0, deg: 90},
		{cx: 0, cy: 0, tx: 5, ty: 5, deg: 135},
		{cx: 0, cy: 0, tx: 0, ty: 5, deg: 180},
		{cx: 0, cy: 0, tx: -5, ty: -5, deg: 315},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test#%v", i), func(t *testing.T) {
			d := angle(tc.cx, tc.cy, tc.tx, tc.ty)
			common.Assert(t, d == tc.deg, "Angle mismatch, got %v, but expected %v", d, tc.deg)
		})
	}
}

func TestExample1(t *testing.T) {
	// 5 X 5, Best @3,4 = 8
	check(t, 5, 5, example1, 10, 3, 4, 8)
}

func TestExample2(t *testing.T) {
	// 10 X 10, Best @5,8 = 33
	check(t, 10, 10, example2, 40, 5, 8, 33)
}

func TestExample3(t *testing.T) {
	// 10 X 10, Best @1,2 = 35
	check(t, 10, 10, example3, 40, 1, 2, 35)
}

func TestExample4(t *testing.T) {
	// 10 X 10, Best @6,3 = 41
	check(t, 10, 10, example4, 50, 6, 3, 41)
}

func TestExample5(t *testing.T) {
	// 20 X 20, Best @11,13 = 210
	check(t, 20, 20, example5, 300, 11, 13, 210)
}

func TestInput(t *testing.T) {
	m, e := newMapFromFile(42, 42, "input")
	common.Assert(t, e == nil, "Failed to create map: %v", e)
	common.Assert(t, m != nil, "Returned nil")
	checkMap(t, m, 410, 26, 36, 347)
}

func TestVaporize(t *testing.T) {
	m, e := newMapFromFile(42, 42, "input")
	common.Assert(t, e == nil, "Failed to create map: %v", e)
	common.Assert(t, m != nil, "Returned nil")
	s := m.asteroids[36*m.width+26]
	common.Assert(t, s != nil, "No asteroid at 26 X 36 to place the monitoring station on")
	vo := m.vaporizeAsteroidsFrom(s)
	common.Assert(t, len(vo) == 409, "Less then 200 asteroids vaporized")
	ta := vo[199]
	common.Assert(t, ta.mapX == 8, "Target X mismatch, expected 8 but got %v", ta.mapX)
	common.Assert(t, ta.mapY == 29, "Target Y mismatch, expected 29 but got %v", ta.mapY)
	a := m.part2Answer()
	common.Assert(t, a == 829, "Part 2 answer mismatch, expected 829 but got %v", a)
}
