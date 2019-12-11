package intcodecpu

import (
	"fmt"
	"testing"

	"aoc2019/common"
)

func TestOpCodeToString(t *testing.T) {
	common.Assert(t, opERR.String() == "ERR", "opERR did not return expected string")
	common.Assert(t, opADD.String() == "ADD", "opADD did not return expected string")
	common.Assert(t, opMUL.String() == "MUL", "opMUL did not return expected string")
	common.Assert(t, opINP.String() == "INP", "opINP did not return expected string")
	common.Assert(t, opOUT.String() == "OUT", "opOUT did not return expected string")
	common.Assert(t, opJMT.String() == "JMT", "opJMT did not return expected string")
	common.Assert(t, opJMF.String() == "JMF", "opJMF did not return expected string")
	common.Assert(t, opLES.String() == "LES", "opLES did not return expected string")
	common.Assert(t, opEQU.String() == "EQU", "opEQU did not return expected string")
	common.Assert(t, opCRB.String() == "CRB", "opCRB did not return expected string")
	common.Assert(t, opHLT.String() == "HLT", "opHLT did not return expected string")
	common.Assert(t, opCode(42).String() == "Unknown OpCode", "Invalid opCode did not return expected string")
}

func executeAndCheckResult(t *testing.T, cpu *CPU, epc uint) {
	done, err := cpu.executeOpCode()
	common.Assert(t, err == nil, "Test failed with an error: %v", err)
	common.Assert(t, done == false, "Test failed with the done flag set")
	common.Assert(t, cpu.pc == epc, "Test failed with unexpected pc: %v, excpected %v", cpu.pc, epc)
}

func TestOpADD(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected int
	}{
		{base: 2, code: []int{1, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 42},     // pos + pos = 21 + 21
		{base: 2, code: []int{101, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 26},   // im + pos = 5 + 21
		{base: 2, code: []int{201, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 31},   // rel + pos = 10 + 21
		{base: 2, code: []int{1001, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 27},  // pos + im = 21 + 6
		{base: 2, code: []int{2001, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 33},  // pos + rel = 21 + 12
		{base: 2, code: []int{1101, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 11},  // im + im = 5 + 6
		{base: 2, code: []int{2201, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 22},  // rel + rel = 10 + 12
		{base: 2, code: []int{22201, 5, 6, 2, 99, 21, 21, 10, 12}, expected: 22}, // rel + rel = 10 + 12 => rel
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, 4)
			if tc.code[0]/10000 == 0 {
				common.Assert(t, cpu.memory[2] == tc.expected, "Test failed for %v with unexpected result %v, excpected %v", tc.code[0], cpu.memory[2], tc.expected)
			} else {
				common.Assert(t, cpu.memory[4] == tc.expected, "Test failed for %v with unexpected result %v, excpected %v", tc.code[0], cpu.memory[4], tc.expected)
			}
		})
	}
}

func TestOpCRB(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected int
	}{
		{base: 5, code: []int{9, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: 7},    // base += pos
		{base: 5, code: []int{109, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: 8},  // base += im
		{base: 5, code: []int{209, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: 12}, // base += rel [base + pos[base + 3] == 5 + 7]
		{base: 5, code: []int{109, -5}, expected: 0},                            // base += im
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, 2)
			common.Assert(t, cpu.relBase == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.relBase, tc.expected)
		})
	}
}

func TestOpEQU(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected int
	}{
		{base: 2, code: []int{8, 4, 5, 0, 42, 42, -1, -2}, expected: 1},    // pos == pos = 42 == 42
		{base: 2, code: []int{8, 4, 5, 0, 42, 44, -1, -1}, expected: 0},    // pos == pos = 42 == 44
		{base: 2, code: []int{108, 4, 5, 0, -1, 4, -2, -3}, expected: 1},   // im  == pos =  4 ==  4
		{base: 2, code: []int{108, 4, 5, 0, -1, 5, -1, -1}, expected: 0},   // im  == pos =  4 ==  5
		{base: 2, code: []int{208, 4, 5, 0, -1, 22, 22, -2}, expected: 1},  // rel == pos = 22 == 22
		{base: 2, code: []int{208, 4, 5, 0, -1, 22, 23, -1}, expected: 0},  // rel == pos = 23 == 22
		{base: 2, code: []int{1008, 4, 5, 0, 5, -1, -2, -3}, expected: 1},  // pos == im  =  5 ==  5
		{base: 2, code: []int{1008, 4, 5, 0, 2, -1, -1, -1}, expected: 0},  // pos == im  =  2 ==  5
		{base: 2, code: []int{2008, 4, 5, 0, 2, -1, -2, 2}, expected: 1},   // pos == rel =  2 ==  2
		{base: 2, code: []int{2008, 4, 5, 0, 2, -1, -2, 3}, expected: 0},   // pos == rel =  2 ==  3
		{base: 2, code: []int{1108, 7, 7, 0, -1, -2, -3, -4}, expected: 1}, // im  == im  =  7 ==  7
		{base: 2, code: []int{1108, 6, 8, 0, -1, -1, -1, -1}, expected: 0}, // im  == im  =  6 ==  8
		{base: 2, code: []int{2208, 4, 5, 0, -1, -2, 8, 8}, expected: 1},   // rel == rel =  8 ==  8
		{base: 2, code: []int{2208, 4, 5, 0, -1, -1, 8, 9}, expected: 0},   // rel == rel =  8 ==  9
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpHLT(t *testing.T) {
	cpu := &CPU{}
	cpu.memory = []int{99}
	done, err := cpu.executeOpCode()
	common.Assert(t, err == nil, "Test failed with an error: %v", err)
	common.Assert(t, done == true, "Test failed with the done flag set")
	common.Assert(t, cpu.pc == 0, "Test failed with unexpected pc: %v, excpected 0", cpu.pc)
}

func TestOpINP(t *testing.T) {
	cpu := &CPU{memory: []int{3, 0}, stdIn: func() int { return 42 }}
	executeAndCheckResult(t, cpu, 2)
	common.Assert(t, cpu.memory[0] == 42, "Test failed with unexpected result %v, excpected 42", cpu.memory[0])
}

func TestOpJMF(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected uint
	}{
		{base: 2, code: []int{6, 3, 4, 0, 42}, expected: 42},   // pos, pos => 0, 42, jump
		{base: 2, code: []int{6, 3, 4, 1, 42}, expected: 3},    // pos, pos => 1, 42, no jump
		{base: 2, code: []int{106, 0, 4, 0, 42}, expected: 42}, // im,  pos => 0, 42, jump
		{base: 2, code: []int{106, 3, 4, 0, 42}, expected: 3},  // im,  pos => 3, 42, no jump
		{base: 2, code: []int{1006, 3, 4, 0, 42}, expected: 4}, // pos, im  => 0,  4, jump
		{base: 2, code: []int{1006, 3, 4, 1, 42}, expected: 3}, // pos, im  => 1,  4, no jump
		{base: 2, code: []int{1106, 0, 4, 0, 42}, expected: 4}, // im,  im  => 0,  4, jump
		{base: 2, code: []int{1106, 3, 4, 0, 42}, expected: 3}, // im,  im  => 3,  4, no jump
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, tc.expected)
		})
	}
}

func TestOpJMT(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected uint
	}{
		{base: 2, code: []int{5, 3, 4, 0, 42}, expected: 3},    // pos, pos => 0, 42, no jump
		{base: 2, code: []int{5, 3, 4, 1, 42}, expected: 42},   // pos, pos => 1, 42, jump
		{base: 2, code: []int{105, 0, 4, 0, 42}, expected: 3},  // im,  pos => 0, 42, no jump
		{base: 2, code: []int{105, 3, 4, 0, 42}, expected: 42}, // im,  pos => 3, 42, jump
		{base: 2, code: []int{1005, 3, 4, 0, 42}, expected: 3}, // pos, im  => 0,  4, no jump
		{base: 2, code: []int{1005, 3, 4, 1, 42}, expected: 4}, // pos, im  => 1,  4, jump
		{base: 2, code: []int{1105, 0, 4, 0, 42}, expected: 3}, // im,  im  => 0,  4, no jump
		{base: 2, code: []int{1105, 3, 4, 0, 42}, expected: 4}, // im,  im  => 3,  4, jump
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, tc.expected)
		})
	}
}

func TestOpLES(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected int
	}{
		{base: 2, code: []int{7, 4, 5, 0, 1, 42}, expected: 1},     // pos < pos =  1 < 42
		{base: 2, code: []int{7, 4, 5, 0, 42, 42}, expected: 0},    // pos < pos = 42 < 42
		{base: 2, code: []int{107, 4, 5, 0, 1, 42}, expected: 1},   // im  < pos =  4 < 42
		{base: 2, code: []int{107, 4, 5, 0, 1, 4}, expected: 0},    // im  < pos =  4 <  4
		{base: 2, code: []int{1007, 4, 5, 0, 1, 42}, expected: 1},  // pos < im  =  1 <  5
		{base: 2, code: []int{1007, 4, 5, 0, 42, 42}, expected: 0}, // pos < im  = 42 <  5
		{base: 2, code: []int{1107, 4, 5, 0, 1, 42}, expected: 1},  // im  < im  =  4 <  5
		{base: 2, code: []int{1107, 4, 3, 0, 1, 42}, expected: 0},  // im  < im  =  4 <  3
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpMUL(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected int
	}{
		{base: 2, code: []int{2, 5, 6, 0, 99, 1, 42}, expected: 42},     // pos * pos = 1 * 42
		{base: 2, code: []int{102, 5, 6, 0, 99, 1, 42}, expected: 210},  // im * pos = 5 * 42
		{base: 2, code: []int{1002, 5, 6, 0, 99, 1, 42}, expected: 6},   // pos * im = 1 * 6
		{base: 2, code: []int{1102, 5, 6, 0, 99, 1, 42}, expected: 30},  // im * im = 5 * 6
		{base: 2, code: []int{10002, 5, 6, 0, 99, 1, 42}, expected: 42}, // pos * pos = 1 * 42
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpOUT(t *testing.T) {
	tt := []struct {
		base     int
		code     []int
		expected int
	}{
		{base: 2, code: []int{4, 0}, expected: 4},         // pos -> out
		{base: 2, code: []int{104, 0}, expected: 0},       // im -> out
		{base: 2, code: []int{1004, 0}, expected: 1004},   // pos -> out
		{base: 2, code: []int{1104, 0}, expected: 0},      // im -> out
		{base: 2, code: []int{10004, 0}, expected: 10004}, // pos -> out
	}
	out := -1
	cpu := &CPU{memory: []int{3, 0}, stdOut: func(data int) { out = data }}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			cpu.relBase = tc.base
			executeAndCheckResult(t, cpu, 2)
			common.Assert(t, out == tc.expected, "Test failed with unexpected result %v, excpected %v", out, tc.expected)
		})
	}
}
