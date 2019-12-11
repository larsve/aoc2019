package main

import (
	"testing"

	"aoc2019/common"
	cpu "aoc2019/common/intcodecpu"
)

func TestDay2Part1(t *testing.T) {
	expected := 3790689
	v, e := part1()
	common.Assert(t, e == nil, "Part1 ended with error %v", e)
	common.Assert(t, v == expected, "Part1 did not give the correct answer %v, got %b", expected, v)
}

func TestDay2Part2(t *testing.T) {
	expected := 6533
	v, e := part2()
	common.Assert(t, e == nil, "Part2 ended with error %v", e)
	common.Assert(t, v == expected, "Part2 did not give the correct answer %v, got %b", expected, v)
}

func BenchmarkDay2Program(b *testing.B) {
	c := cpu.NewProgram(gravityAssistProgram)
	for n := 0; n < b.N; n++ {
		c.Reset()
		e := c.Run()
		common.Assert(b, e == nil, "Iteration %v ended with an error: %v", n, e)
	}
}
