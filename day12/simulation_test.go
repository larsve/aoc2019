package main

import (
	"fmt"
	"testing"

	"aoc2019/common"
)

func example1Simulation() *simulation {
	sim := &simulation{}
	sim.addMoon("Io", -1, 0, 2)
	sim.addMoon("Europa", 2, -10, -7)
	sim.addMoon("Ganymede", 4, -8, 8)
	sim.addMoon("Callisto", 3, 5, -1)
	return sim
}

func example2Simulation() *simulation {
	sim := &simulation{}
	sim.addMoon("Io", -8, -10, 0)
	sim.addMoon("Europa", 5, 5, 10)
	sim.addMoon("Ganymede", 2, -7, 3)
	sim.addMoon("Callisto", 9, -8, -3)
	return sim
}

func TestExample1Part1(t *testing.T) {
	sim := example1Simulation()

	tt := []struct {
		pos [4]vector
		vel [4]vector
	}{
		{pos: [4]vector{{2, -1, 1}, {3, -7, -4}, {1, -7, 5}, {2, 2, 0}}, vel: [4]vector{{3, -1, -1}, {1, 3, 3}, {-3, 1, -3}, {-1, -3, 1}}},
		{pos: [4]vector{{5, -3, -1}, {1, -2, 2}, {1, -4, -1}, {1, -4, 2}}, vel: [4]vector{{3, -2, -2}, {-2, 5, 6}, {0, 3, -6}, {-1, -6, 2}}},
		{pos: [4]vector{{5, -6, -1}, {0, 0, 6}, {2, 1, -5}, {1, -8, 2}}, vel: [4]vector{{0, -3, 0}, {-1, 2, 4}, {1, 5, -4}, {0, -4, 0}}},
		{pos: [4]vector{{2, -8, 0}, {2, 1, 7}, {2, 3, -6}, {2, -9, 1}}, vel: [4]vector{{-3, -2, 1}, {2, 1, 1}, {0, 2, -1}, {1, -1, -1}}},
		{pos: [4]vector{{-1, -9, 2}, {4, 1, 5}, {2, 2, -4}, {3, -7, -1}}, vel: [4]vector{{-3, -1, 2}, {2, 0, -2}, {0, -1, 2}, {1, 2, -2}}},
		{pos: [4]vector{{-1, -7, 3}, {3, 0, 0}, {3, -2, 1}, {3, -4, -2}}, vel: [4]vector{{0, 2, 1}, {-1, -1, -5}, {1, -4, 5}, {0, 3, -1}}},
		{pos: [4]vector{{2, -2, 1}, {1, -4, -4}, {3, -7, 5}, {2, 0, 0}}, vel: [4]vector{{3, 5, -2}, {-2, -4, -4}, {0, -5, 4}, {-1, 4, 2}}},
		{pos: [4]vector{{5, 2, -2}, {2, -7, -5}, {0, -9, 6}, {1, 1, 3}}, vel: [4]vector{{3, 4, -3}, {1, -3, -1}, {-3, -2, 1}, {-1, 1, 3}}},
		{pos: [4]vector{{5, 3, -4}, {2, -9, -3}, {0, -8, 4}, {1, 1, 5}}, vel: [4]vector{{0, 1, -2}, {0, -2, 2}, {0, 1, -2}, {0, 0, 2}}},
		{pos: [4]vector{{2, 1, -3}, {1, -8, 0}, {3, -6, 1}, {2, 0, 4}}, vel: [4]vector{{-3, -2, 1}, {-1, 1, 3}, {3, 2, -3}, {1, -1, -1}}},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test#%v", i), func(t *testing.T) {
			sim.step()
			for x, m := range sim.moons {
				common.Assert(t, m.vel == tc.vel[x], "Not the expected velocity for %v, expected %v but got %v", m.name, tc.vel[x], m.vel)
				common.Assert(t, m.pos == tc.pos[x], "Not the expected position for %v, expected %v but got %v", m.name, tc.pos[x], m.pos)
			}
		})
	}
	e := sim.energy()
	common.Assert(t, e == 179, "Not the expected total energy, expected %v but got %v", 179, e)
}

func TestExample1Part2(t *testing.T) {
	sim := example1Simulation()
	steps := sim.loop()
	common.Assert(t, steps == 2772, "Not the expected step count between repeating patterns, expected %v but got %v", 2772, steps)
	common.Assert(t, sim.stepCount == 44, "Not the expected step count, expected %v but got %v", 44, sim.stepCount)
}

func TestExample2Part1(t *testing.T) {
	sim := example2Simulation()

	tt := []struct {
		pos [4]vector
		vel [4]vector
	}{
		{pos: [4]vector{{-9, -10, 1}, {4, 10, 9}, {8, -10, -3}, {5, -10, 3}}, vel: [4]vector{{-2, -2, -1}, {-3, 7, -2}, {5, -1, -2}, {0, -4, 5}}},
		{pos: [4]vector{{-10, 3, -4}, {5, -25, 6}, {13, 1, 1}, {0, 1, 7}}, vel: [4]vector{{-5, 2, 0}, {1, 1, -4}, {5, -2, 2}, {-1, -1, 2}}},
		{pos: [4]vector{{15, -6, -9}, {-4, -11, 3}, {0, -1, 11}, {-3, -2, 5}}, vel: [4]vector{{-5, 4, 0}, {-3, -10, 0}, {7, 4, 3}, {1, 2, -3}}},
		{pos: [4]vector{{14, -12, -4}, {-1, 18, 8}, {-5, -14, 8}, {0, -12, -2}}, vel: [4]vector{{11, 3, 0}, {-5, 2, 3}, {1, -2, 0}, {-7, -3, -3}}},
		{pos: [4]vector{{-23, 4, 1}, {20, -31, 13}, {-4, 6, 1}, {15, 1, -5}}, vel: [4]vector{{-7, -1, 2}, {5, 3, 4}, {-1, 1, -3}, {3, -3, -3}}},
		{pos: [4]vector{{36, -10, 6}, {-18, 10, 9}, {8, -12, -3}, {-18, -8, -2}}, vel: [4]vector{{5, 0, 3}, {-3, -7, 5}, {-2, 1, -7}, {0, 6, -1}}},
		{pos: [4]vector{{-33, -6, 5}, {13, -9, 2}, {11, -8, 2}, {17, 3, 1}}, vel: [4]vector{{-5, -4, 7}, {-2, 11, 3}, {8, -6, -7}, {-1, -1, -3}}},
		{pos: [4]vector{{30, -8, 3}, {-2, -4, 0}, {-18, -7, 15}, {-2, -1, -8}}, vel: [4]vector{{3, 3, 0}, {4, -13, 2}, {-8, 2, -2}, {1, 8, 0}}},
		{pos: [4]vector{{-25, -1, 4}, {2, -9, 0}, {32, -8, 14}, {-1, -2, -8}}, vel: [4]vector{{1, -3, 4}, {-3, 13, -1}, {5, -4, 6}, {-3, -6, -9}}},
		{pos: [4]vector{{8, -12, -9}, {13, 16, -3}, {-29, -11, -1}, {16, -13, 23}}, vel: [4]vector{{-7, 3, 0}, {3, -11, -5}, {-3, 7, 4}, {7, 1, 1}}},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test#%v", i), func(t *testing.T) {
			sim.steps(10)
			for x, m := range sim.moons {
				common.Assert(t, m.vel == tc.vel[x], "Not the expected velocity for %v, expected %v but got %v", m.name, tc.vel[x], m.vel)
				common.Assert(t, m.pos == tc.pos[x], "Not the expected position for %v, expected %v but got %v", m.name, tc.pos[x], m.pos)
			}
		})
	}
	e := sim.energy()
	common.Assert(t, e == 1940, "Not the expected total energy, expected %v but got %v", 1940, e)
}

func TestPart1(t *testing.T) {
	sim := part1()
	e := sim.energy()
	common.Assert(t, e == 14780, "Not the expected energy, expected %v but got %v", 14780, e)
}

func TestPart2(t *testing.T) {
	sim, steps := part2()
	common.Assert(t, steps == 279751820342592, "Not the expected step count, expected %v but got %v", 279751820342592, steps)
	common.Assert(t, sim.stepCount == 286332, "Not the expected step count, expected %v but got %v", 286332, sim.stepCount)
}
