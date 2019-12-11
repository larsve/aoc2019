package intcodecpu

import (
	"fmt"
	"reflect"
	"testing"

	"aoc2019/common"
)

func TestGetOpCode(t *testing.T) {
	tt := []struct {
		codes          []int
		expectedOpCode opCode
	}{
		{codes: []int{1, 101, 1001, 10001, 201, 2001, 20001}, expectedOpCode: opADD},
		{codes: []int{2, 102, 1002, 10002, 202, 2002, 20002}, expectedOpCode: opMUL},
		{codes: []int{3, 103, 1003, 10003, 203, 2003, 20003}, expectedOpCode: opINP},
		{codes: []int{4, 104, 1004, 10004, 204, 2004, 20004}, expectedOpCode: opOUT},
		{codes: []int{5, 105, 1005, 10005, 205, 2005, 20005}, expectedOpCode: opJMT},
		{codes: []int{6, 106, 1006, 10006, 206, 2006, 20006}, expectedOpCode: opJMF},
		{codes: []int{7, 107, 1007, 10007, 207, 2007, 20007}, expectedOpCode: opLES},
		{codes: []int{8, 108, 1008, 10008, 208, 2008, 20008}, expectedOpCode: opEQU},
		{codes: []int{9, 109, 1009, 10009, 209, 2009, 20009}, expectedOpCode: opCRB},
		{codes: []int{99}, expectedOpCode: opHLT},
		{codes: []int{-1, 42, 142, 1242, 12342}, expectedOpCode: opERR},
	}
	cpu := &CPU{memory: []int{0}}
	for _, tc := range tt {
		t.Run(tc.expectedOpCode.String(), func(t *testing.T) {
			for _, code := range tc.codes {
				cpu.memory[0] = code
				opCode := cpu.getOpCode()
				common.Assert(t, opCode == tc.expectedOpCode, "Not the expected opCode for %v, expected %v but got %v", code, tc.expectedOpCode, opCode)
			}
		})
	}
}

func TestGetOpParams(t *testing.T) {
	c := &CPU{}
	test := func(p [4]int, e [6]int) {
		c.memory = p[:]
		p1m, p1, p2m, p2, p3m, p3 := c.getOpParams(4)
		a := [6]int{p1m, p1, p2m, p2, p3m, p3}
		common.Assert(t, a == e, "Params %v did not generate the expected output: %v, got: %v", p, e, a)
	}
	test([4]int{0, 1, 2, 3}, [6]int{0, 1, 0, 2, 0, 3})
	test([4]int{100, 10, 20, 30}, [6]int{1, 10, 0, 20, 0, 30})
	test([4]int{1000, 15, 30, 45}, [6]int{0, 15, 1, 30, 0, 45})
	test([4]int{10000, 20, 40, 60}, [6]int{0, 20, 0, 40, 1, 60})
	test([4]int{12300, -1, -1, -1}, [6]int{3, -1, 2, -1, 1, -1})
}

func TestCpuLda(t *testing.T) {
	cpu := &CPU{memory: []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9, 111, 222, 333, 444, 555, 666, 777, 888, 999}}
	tt := []struct {
		param int
		base  int
		value [3]int // Expected values for mode 0 (pos), 1 (im) and 2 (rel)
	}{
		{param: 0, base: 10, value: [3]int{0, 0, 111}},
		{param: 5, base: 10, value: [3]int{1, 5, 666}},
		{param: 8, base: 10, value: [3]int{7, 8, 999}},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.relBase = tc.base
			for mode := 0; mode < 3; mode++ {
				v := cpu.lda(tc.param, mode)
				common.Assert(t, v == tc.value[mode], "Not the expected parameter value for mode %v, expected %v but got %v", mode, tc.value[mode], v)
			}
		})
	}
}

func TestReset(t *testing.T) {
	p := faultyProgram()
	p.memory[2] = 42
	p.Reset()
	common.Assert(t, p.memory[2] == 0, "Reset failed")
}

func TestRun(t *testing.T) {
	p := smallAddProgram()
	e := p.Run()
	common.Assert(t, e == nil, "Run ended with an error: %v", e)
	m := p.DumpMemory(0, 1)
	common.Assert(t, m[0] == 2, "Run didn't give the expected result, expected 2 but got %v", p.memory[0])

	p = faultyProgram()
	e = p.Run()
	common.Assert(t, e != nil, "Run didn't end with an error")
}

func TestRunDay5Examples(t *testing.T) {
	// Part 1
	out := -1
	p := day5Example1Program()
	p.stdIn = func() int { return 42 }
	p.stdOut = func(data int) { out = data }
	e := p.Run()
	common.Assert(t, e == nil, "Example1 run ended with an error: %v", e)
	common.Assert(t, out == 42, "Example1 failed with wrong output, got %v", out)

	p = day5Example2Program()
	e = p.Run()
	common.Assert(t, e == nil, "Example2 run ended with an error: %v", e)

	p = day5Example3Program()
	e = p.Run()
	common.Assert(t, e == nil, "Example3 run ended with an error: %v", e)

	// Part 2
	c := day5Example4Program()
	c.stdIn = func() int { return 8 }
	c.stdOut = func(data int) { out = data }
	e = c.Run()
	common.Assert(t, e == nil, "Example4_1 falied with error: %v", e)
	common.Assert(t, out == 1000, "Example4_1 output error, got: %v", out)

	c.Reset()
	c.stdIn = func() int { return 0 }
	//c.DebugOutput = func(msg string) { fmt.Print(msg) }
	e = c.Run()
	common.Assert(t, e == nil, "Example4_2 falied with error: %v", e)
	common.Assert(t, out == 999, "Example4_2 output error, got: %v", out)

	c.Reset()
	c.stdIn = func() int { return 9 }
	e = c.Run()
	common.Assert(t, e == nil, "Example4_3 falied with error: %v", e)
	common.Assert(t, out == 1001, "Example4_3 output error, got: %v", out)
}

func BenchmarkSmallAddProgram(b *testing.B) {
	c := smallAddProgram()
	for n := 0; n < b.N; n++ {
		c.Reset()
		e := c.Run()
		common.Assert(b, e == nil, "Iteration %v ended with an error: %v", n, e)
	}
}

func TestDay9Examples(t *testing.T) {
	run := func(code []int) (output []int, err error) {
		output = []int{}
		c := NewProgram(code)
		c.stdOut = func(data int) { output = append(output, data) }
		c.DebugOutput = func(msg string) { fmt.Print(msg) }
		err = c.Run()
		return
	}
	// Should generate a copy of itself
	code := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	output, e := run(code)
	common.Assert(t, e == nil, "TestDay9Example1 falied with error: %v", e)
	common.Assert(t, len(output) == 16, "TestDay9Example1 did not generate expected output: %v", output)
	common.Assert(t, reflect.DeepEqual(code, output), "TestDay9Example1 did not generate expected output: %v", output)

	// Should generate a 16-digit number
	output, e = run([]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0})
	common.Assert(t, e == nil, "TestDay9Example2 falied with error: %v", e)
	common.Assert(t, len(output) == 1, "TestDay9Example2 did not generate expected output: %v", output)
	common.Assert(t, output[0] == 1219070632396864, "TestDay9Example2 did not generate expected output: %v", output[0])

	// Should generate 1125899906842624 as output
	output, e = run([]int{104, 1125899906842624, 99})
	common.Assert(t, e == nil, "TestDay9Example3 falied with error: %v", e)
	common.Assert(t, len(output) == 1, "TestDay9Example3 did not generate expected output: %v", output)
	common.Assert(t, output[0] == 1125899906842624, "TestDay9Example3 did not generate expected output: %v", output[0])
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
