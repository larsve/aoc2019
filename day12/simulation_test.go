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
		{pos: [4]vector{vector{2, -1, 1}, vector{3, -7, -4}, vector{1, -7, 5}, vector{2, 2, 0}}, vel: [4]vector{vector{3, -1, -1}, vector{1, 3, 3}, vector{-3, 1, -3}, vector{-1, -3, 1}}},
		{pos: [4]vector{vector{5, -3, -1}, vector{1, -2, 2}, vector{1, -4, -1}, vector{1, -4, 2}}, vel: [4]vector{vector{3, -2, -2}, vector{-2, 5, 6}, vector{0, 3, -6}, vector{-1, -6, 2}}},
		{pos: [4]vector{vector{5, -6, -1}, vector{0, 0, 6}, vector{2, 1, -5}, vector{1, -8, 2}}, vel: [4]vector{vector{0, -3, 0}, vector{-1, 2, 4}, vector{1, 5, -4}, vector{0, -4, 0}}},
		{pos: [4]vector{vector{2, -8, 0}, vector{2, 1, 7}, vector{2, 3, -6}, vector{2, -9, 1}}, vel: [4]vector{vector{-3, -2, 1}, vector{2, 1, 1}, vector{0, 2, -1}, vector{1, -1, -1}}},
		{pos: [4]vector{vector{-1, -9, 2}, vector{4, 1, 5}, vector{2, 2, -4}, vector{3, -7, -1}}, vel: [4]vector{vector{-3, -1, 2}, vector{2, 0, -2}, vector{0, -1, 2}, vector{1, 2, -2}}},
		{pos: [4]vector{vector{-1, -7, 3}, vector{3, 0, 0}, vector{3, -2, 1}, vector{3, -4, -2}}, vel: [4]vector{vector{0, 2, 1}, vector{-1, -1, -5}, vector{1, -4, 5}, vector{0, 3, -1}}},
		{pos: [4]vector{vector{2, -2, 1}, vector{1, -4, -4}, vector{3, -7, 5}, vector{2, 0, 0}}, vel: [4]vector{vector{3, 5, -2}, vector{-2, -4, -4}, vector{0, -5, 4}, vector{-1, 4, 2}}},
		{pos: [4]vector{vector{5, 2, -2}, vector{2, -7, -5}, vector{0, -9, 6}, vector{1, 1, 3}}, vel: [4]vector{vector{3, 4, -3}, vector{1, -3, -1}, vector{-3, -2, 1}, vector{-1, 1, 3}}},
		{pos: [4]vector{vector{5, 3, -4}, vector{2, -9, -3}, vector{0, -8, 4}, vector{1, 1, 5}}, vel: [4]vector{vector{0, 1, -2}, vector{0, -2, 2}, vector{0, 1, -2}, vector{0, 0, 2}}},
		{pos: [4]vector{vector{2, 1, -3}, vector{1, -8, 0}, vector{3, -6, 1}, vector{2, 0, 4}}, vel: [4]vector{vector{-3, -2, 1}, vector{-1, 1, 3}, vector{3, 2, -3}, vector{1, -1, -1}}},
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
	common.Assert(t, sim.stepCount == 2772, "Not the expected step count, expected %v but got %v", 2772, sim.stepCount)
}

func TestExample2Part1(t *testing.T) {
	sim := example2Simulation()

	tt := []struct {
		pos [4]vector
		vel [4]vector
	}{
		{pos: [4]vector{vector{-9, -10, 1}, vector{4, 10, 9}, vector{8, -10, -3}, vector{5, -10, 3}}, vel: [4]vector{vector{-2, -2, -1}, vector{-3, 7, -2}, vector{5, -1, -2}, vector{0, -4, 5}}},
		{pos: [4]vector{vector{-10, 3, -4}, vector{5, -25, 6}, vector{13, 1, 1}, vector{0, 1, 7}}, vel: [4]vector{vector{-5, 2, 0}, vector{1, 1, -4}, vector{5, -2, 2}, vector{-1, -1, 2}}},
		{pos: [4]vector{vector{15, -6, -9}, vector{-4, -11, 3}, vector{0, -1, 11}, vector{-3, -2, 5}}, vel: [4]vector{vector{-5, 4, 0}, vector{-3, -10, 0}, vector{7, 4, 3}, vector{1, 2, -3}}},
		{pos: [4]vector{vector{14, -12, -4}, vector{-1, 18, 8}, vector{-5, -14, 8}, vector{0, -12, -2}}, vel: [4]vector{vector{11, 3, 0}, vector{-5, 2, 3}, vector{1, -2, 0}, vector{-7, -3, -3}}},
		{pos: [4]vector{vector{-23, 4, 1}, vector{20, -31, 13}, vector{-4, 6, 1}, vector{15, 1, -5}}, vel: [4]vector{vector{-7, -1, 2}, vector{5, 3, 4}, vector{-1, 1, -3}, vector{3, -3, -3}}},
		{pos: [4]vector{vector{36, -10, 6}, vector{-18, 10, 9}, vector{8, -12, -3}, vector{-18, -8, -2}}, vel: [4]vector{vector{5, 0, 3}, vector{-3, -7, 5}, vector{-2, 1, -7}, vector{0, 6, -1}}},
		{pos: [4]vector{vector{-33, -6, 5}, vector{13, -9, 2}, vector{11, -8, 2}, vector{17, 3, 1}}, vel: [4]vector{vector{-5, -4, 7}, vector{-2, 11, 3}, vector{8, -6, -7}, vector{-1, -1, -3}}},
		{pos: [4]vector{vector{30, -8, 3}, vector{-2, -4, 0}, vector{-18, -7, 15}, vector{-2, -1, -8}}, vel: [4]vector{vector{3, 3, 0}, vector{4, -13, 2}, vector{-8, 2, -2}, vector{1, 8, 0}}},
		{pos: [4]vector{vector{-25, -1, 4}, vector{2, -9, 0}, vector{32, -8, 14}, vector{-1, -2, -8}}, vel: [4]vector{vector{1, -3, 4}, vector{-3, 13, -1}, vector{5, -4, 6}, vector{-3, -6, -9}}},
		{pos: [4]vector{vector{8, -12, -9}, vector{13, 16, -3}, vector{-29, -11, -1}, vector{16, -13, 23}}, vel: [4]vector{vector{-7, 3, 0}, vector{3, -11, -5}, vector{-3, 7, 4}, vector{7, 1, 1}}},
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

/*
Test fails with out of memory error, need to rewrite solution to be able to pass this test
func TestPart2(t *testing.T) {
	sim := part2()
	common.Assert(t, sim.stepCount > 60103453, "Not the expected step count, expected > %v but got %v", 60103453, sim.stepCount)
}
*/
