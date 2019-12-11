package main

import (
	cpu "aoc2019/common/intcodecpu"
)

var gravityAssistProgram = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}

func part1() (int, error) {
	c := cpu.NewProgram(gravityAssistProgram)
	// Restore the "1202 program alert code"
	c.PatchMemory(1, []int{12, 2})
	if e := c.Run(); e != nil {
		return -1, e
	}
	mem := c.DumpMemory(0, 1)
	return mem[0], nil
}

func part2() (int, error) {
	c := cpu.NewProgram(gravityAssistProgram)
	for i := 0; i <= 9999; i++ {
		c.Reset()
		c.PatchMemory(1, []int{i / 100, i % 100})
		if e := c.Run(); e != nil {
			return -1, e
		}
		mem := c.DumpMemory(0, 1)
		if mem[0] == 19690720 {
			return i, nil
		}
	}
	return -1, nil
}
