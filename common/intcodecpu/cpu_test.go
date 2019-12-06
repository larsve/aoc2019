package intcodecpu

import (
	"testing"

	"aoc2019/common"
)

func TestReset(t *testing.T) {
	p := faultyProgram()
	p.memory[2] = 42
	p.reset()
	common.Assert(t, p.memory[2] == 0, "Reset failed")
}

func TestRun(t *testing.T) {
	p := smallAddProgram()
	r, e := p.Run()
	common.Assert(t, e == nil, "Run ended with an error: %v", e)
	common.Assert(t, r == 2, "Run didn't give the expected result, expected 2 but got %v", r)

	p = faultyProgram()
	_, e = p.Run()
	common.Assert(t, e != nil, "Run didn't end with an error")
}

func TestRunDay5Examples(t *testing.T) {
	out := -1
	p := day5Example1Program()
	p.stdIn = func() int { return 42 }
	p.stdOut = func(data int) { out = data }
	_, e := p.Run()
	common.Assert(t, e == nil, "Example1 run ended with an error: %v", e)
	common.Assert(t, out == 42, "Example1 failed with wrong output, got %v", out)

	p = day5Example2Program()
	_, e = p.Run()
	common.Assert(t, e == nil, "Example2 run ended with an error: %v", e)

	p = day5Example3Program()
	_, e = p.Run()
	common.Assert(t, e == nil, "Example3 run ended with an error: %v", e)
}

func BenchmarkSmallAddProgram(b *testing.B) {
	c := smallAddProgram()
	for n := 0; n < b.N; n++ {
		c.reset()
		_, e := c.Run()
		common.Assert(b, e == nil, "Iteration %v ended with an error: %v", n, e)
	}
}

func BenchmarkDay2Program(b *testing.B) {
	c := day2Program()
	for n := 0; n < b.N; n++ {
		c.reset()
		_, e := c.Run()
		common.Assert(b, e == nil, "Iteration %v ended with an error: %v", n, e)
	}
}

func TestSetAlarmCode(t *testing.T) {
	i := sampleProgram()
	test := func(code, expectedNoun, expectedVerb int) {
		i.SetAlarmCode(code)
		common.Assert(t, i.memory[1] == expectedNoun, "Not the expected noun, expected %v but got %v (alarm code: %v)", expectedNoun, i.memory[1], code)
		common.Assert(t, i.memory[2] == expectedVerb, "Not the expected verb, expected %v but got %v (alarm code: %v)", expectedVerb, i.memory[2], code)
	}
	test(99, 00, 99)
	test(9900, 99, 00)
	test(999, 9, 99)
	test(1000, 10, 0)
}

func TestGetNounAndVerbFor(t *testing.T) {
	c := day2Program()
	common.Assert(t, c.GetNounAndVerbFor(19690720) == 6533, "GetNounAndVerbFor falied")
}

func BenchmarkGetNounAndVerbFor(b *testing.B) {
	c := day2Program()
	for n := 0; n < b.N; n++ {
		c.GetNounAndVerbFor(19690720)
	}
}

func TestDay5Part1(t *testing.T) {
	out := -1
	nzout := 0
	c := day5Program()
	c.stdIn = func() int { return 1 }
	c.stdOut = func(data int) {
		if data != 0 {
			nzout++
		}
		out = data
	}
	_, e := c.Run()
	common.Assert(t, e == nil, "TestDay5Part1 falied with error: %v", e)
	common.Assert(t, nzout == 1, "TestDay5Part1 should have onle one non-zero output, got: %v", nzout)
	common.Assert(t, out == 9219874, "TestDay5Part1 output error, got: %v", out)
}

func TestDay5Part2Example(t *testing.T) {
	out := -1
	c := day5Example4Program()
	c.stdIn = func() int { return 8 }
	c.stdOut = func(data int) { out = data }
	_, e := c.Run()
	common.Assert(t, e == nil, "TestDay5Part1 falied with error: %v", e)
	common.Assert(t, out == 1000, "TestDay5Part1 output error, got: %v", out)

	c.reset()
	c.stdIn = func() int { return 0 }
	_, e = c.Run()
	common.Assert(t, e == nil, "TestDay5Part1 falied with error: %v", e)
	common.Assert(t, out == 999, "TestDay5Part1 output error, got: %v", out)

	c.reset()
	c.stdIn = func() int { return 9 }
	_, e = c.Run()
	common.Assert(t, e == nil, "TestDay5Part1 falied with error: %v", e)
	common.Assert(t, out == 1001, "TestDay5Part1 output error, got: %v", out)
}

func TestDay5Part2(t *testing.T) {
	out := -1
	c := day5Program()
	c.stdIn = func() int { return 5 }
	c.stdOut = func(data int) { out = data }
	_, e := c.Run()
	common.Assert(t, e == nil, "TestDay5Part1 falied with error: %v", e)
	common.Assert(t, out == 5893654, "TestDay5Part1 output error, got: %v", out)
}

func BenchmarkDay5Part2(b *testing.B) {
	c := day5Program()
	c.stdIn = func() int { return 5 }
	c.stdOut = func(data int) {}
	for n := 0; n < b.N; n++ {
		_, e := c.Run()
		common.Assert(b, e == nil, "TestDay5Part2 falied with error: %v", e)
		c.reset()
	}
}

func sampleProgram() *CPU {
	return NewProgram([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
}
func smallAddProgram() *CPU {
	return NewProgram([]int{1, 0, 0, 0, 99})
}
func faultyProgram() *CPU {
	return NewProgram([]int{0, 0, 0, 0, 42})
}
func day5Example1Program() *CPU {
	return NewProgram([]int{3, 0, 4, 0, 99})
}
func day5Example2Program() *CPU {
	return NewProgram([]int{1002, 4, 3, 4, 33})
}
func day5Example3Program() *CPU {
	return NewProgram([]int{1101, 100, -1, 4, 0})
}
func day5Example4Program() *CPU {
	// Takes a single input.
	// If below 8, output == 999
	// If equal to 8, output == 1000
	// If above 8, output == 1001
	return NewProgram([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99})
}
func day2Program() *CPU {
	return NewProgram([]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0})
}
func day5Program() *CPU {
	return NewProgram([]int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 101, 67, 166, 224, 1001, 224, -110, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 2, 62, 66, 224, 101, -406, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 76, 51, 225, 1101, 51, 29, 225, 1102, 57, 14, 225, 1102, 64, 48, 224, 1001, 224, -3072, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1001, 217, 90, 224, 1001, 224, -101, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1101, 57, 55, 224, 1001, 224, -112, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1102, 5, 62, 225, 1102, 49, 68, 225, 102, 40, 140, 224, 101, -2720, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1101, 92, 43, 225, 1101, 93, 21, 225, 1002, 170, 31, 224, 101, -651, 224, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 223, 224, 223, 1, 136, 57, 224, 1001, 224, -138, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 11, 85, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 329, 1001, 223, 1, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 344, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 101, 1, 223, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 374, 1001, 223, 1, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 434, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 449, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 524, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 554, 1001, 223, 1, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 599, 1001, 223, 1, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 629, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226})
}
