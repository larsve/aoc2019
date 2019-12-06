package intcodecpu

import (
	"fmt"
	"testing"

	"aoc2019/common"
)

func TestOpCodeToString(t *testing.T) {
	common.Assert(t, opAdd.String() == "ADD", "opAdd did not return expected string")
	common.Assert(t, opMul.String() == "MUL", "opMul did not return expected string")
	common.Assert(t, opIn.String() == "IN", "opIn did not return expected string")
	common.Assert(t, opOut.String() == "OUT", "opOut did not return expected string")
	common.Assert(t, opErr.String() == "ERR", "opErr did not return expected string")
	common.Assert(t, opExit.String() == "HALT", "opExit did not return expected string")
	common.Assert(t, opCode(42).String() == "Unknown OpCode", "Invalid opCode did not return expected string")
}

func TestGetOpCode(t *testing.T) {
	tt := []struct {
		codes          []int
		expectedOpCode opCode
	}{
		{codes: []int{1, 101, 1001, 10001}, expectedOpCode: opAdd},
		{codes: []int{2, 102, 1002, 10002}, expectedOpCode: opMul},
		{codes: []int{3, 103, 1003, 10003}, expectedOpCode: opIn},
		{codes: []int{4, 104, 1004, 10004}, expectedOpCode: opOut},
		{codes: []int{99}, expectedOpCode: opExit},
		{codes: []int{-1, 42, 142, 1242, 12342}, expectedOpCode: opErr},
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

/*
func testOp(t *testing.T, o opCode, p []int, elda, esta bool, eaddr, eval int) {
	m := &testMemoryAccessor{}
	io := &testIO{in: 42}
	e := o.execute(p, m, io)
	common.Assert(t, e == nil, "OP %v test failed, terminated with an error: %v", o, e)
	common.Assert(t, m.calledLda == elda, "OP %v test failed, LDA call, expected %v, but were %v", o, elda, m.calledLda)
	common.Assert(t, m.calledSta == esta, "OP %v test failed, STA call, expected %v, but were %v", o, esta, m.calledSta)
	common.Assert(t, m.staAddr == eaddr, "OP %v test failed, STA addr, expected %v, but were %v", o, eaddr, m.staAddr)
	common.Assert(t, m.staValue == eval, "OP %v test failed, STA value, expected %v, but were %v", o, eaddr, m.staValue)
}
*/
func executeAndCheckResult(t *testing.T, cpu *CPU, epc uint) {
	done, err := cpu.executeOpCode()
	common.Assert(t, err == nil, "Test failed with an error: %v", err)
	common.Assert(t, done == false, "Test failed with the done flag set")
	common.Assert(t, cpu.pc == epc, "Test failed with unexpected pc: %v, excpected %v", cpu.pc, epc)
}

func TestOpAdd(t *testing.T) {
	tt := []struct {
		code     []int
		expected int
	}{
		{code: []int{1, 5, 6, 0, 99, 21, 21}, expected: 42},     // pos + pos = 21 + 21
		{code: []int{101, 5, 6, 0, 99, 21, 21}, expected: 26},   // im + pos = 5 + 21
		{code: []int{1001, 5, 6, 0, 99, 21, 21}, expected: 27},  // pos + im = 21 + 6
		{code: []int{1101, 5, 6, 0, 99, 21, 21}, expected: 11},  // im + im = 5 + 6
		{code: []int{10001, 5, 6, 0, 99, 21, 21}, expected: 42}, // pos + pos = 21 + 21
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpEqu(t *testing.T) {
	tt := []struct {
		code     []int
		expected int
	}{
		{code: []int{8, 4, 5, 0, 42, 42}, expected: 1},    // pos == pos = 42 == 42
		{code: []int{8, 4, 5, 0, 42, 44}, expected: 0},    // pos == pos = 42 == 42
		{code: []int{108, 4, 5, 0, 42, 4}, expected: 1},   // im  == pos =  4 ==  4
		{code: []int{108, 4, 5, 0, 42, 42}, expected: 0},  // im  == pos =  4 == 42
		{code: []int{1008, 4, 5, 0, 5, 42}, expected: 1},  // pos == im  =  5 ==  5
		{code: []int{1008, 4, 5, 0, 42, 42}, expected: 0}, // pos == im  = 42 ==  5
		{code: []int{1108, 7, 7, 0, 42, 42}, expected: 1}, // im  == im  =  7 ==  7
		{code: []int{1108, 6, 8, 0, 42, 42}, expected: 0}, // im  == im  =  6 ==  8
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpIn(t *testing.T) {
	cpu := &CPU{memory: []int{3, 0}, stdIn: func() int { return 42 }}
	executeAndCheckResult(t, cpu, 2)
	common.Assert(t, cpu.memory[0] == 42, "Test failed with unexpected result %v, excpected 42", cpu.memory[0])
}

func TestOpJmpF(t *testing.T) {
	tt := []struct {
		code     []int
		expected uint
	}{
		{code: []int{6, 3, 4, 0, 42}, expected: 42},   // pos, pos => 0, 42, jump
		{code: []int{6, 3, 4, 1, 42}, expected: 3},    // pos, pos => 1, 42, no jump
		{code: []int{106, 0, 4, 0, 42}, expected: 42}, // im,  pos => 0, 42, jump
		{code: []int{106, 3, 4, 0, 42}, expected: 3},  // im,  pos => 3, 42, no jump
		{code: []int{1006, 3, 4, 0, 42}, expected: 4}, // pos, im  => 0,  4, jump
		{code: []int{1006, 3, 4, 1, 42}, expected: 3}, // pos, im  => 1,  4, no jump
		{code: []int{1106, 0, 4, 0, 42}, expected: 4}, // im,  im  => 0,  4, jump
		{code: []int{1106, 3, 4, 0, 42}, expected: 3}, // im,  im  => 3,  4, no jump
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, tc.expected)
		})
	}
}

func TestOpJmpT(t *testing.T) {
	tt := []struct {
		code     []int
		expected uint
	}{
		{code: []int{5, 3, 4, 0, 42}, expected: 3},    // pos, pos => 0, 42, no jump
		{code: []int{5, 3, 4, 1, 42}, expected: 42},   // pos, pos => 1, 42, jump
		{code: []int{105, 0, 4, 0, 42}, expected: 3},  // im,  pos => 0, 42, no jump
		{code: []int{105, 3, 4, 0, 42}, expected: 42}, // im,  pos => 3, 42, jump
		{code: []int{1005, 3, 4, 0, 42}, expected: 3}, // pos, im  => 0,  4, no jump
		{code: []int{1005, 3, 4, 1, 42}, expected: 4}, // pos, im  => 1,  4, jump
		{code: []int{1105, 0, 4, 0, 42}, expected: 3}, // im,  im  => 0,  4, no jump
		{code: []int{1105, 3, 4, 0, 42}, expected: 4}, // im,  im  => 3,  4, jump
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, tc.expected)
		})
	}
}

func TestOpLess(t *testing.T) {
	tt := []struct {
		code     []int
		expected int
	}{
		{code: []int{7, 4, 5, 0, 1, 42}, expected: 1},     // pos < pos =  1 < 42
		{code: []int{7, 4, 5, 0, 42, 42}, expected: 0},    // pos < pos = 42 < 42
		{code: []int{107, 4, 5, 0, 1, 42}, expected: 1},   // im  < pos =  4 < 42
		{code: []int{107, 4, 5, 0, 1, 4}, expected: 0},    // im  < pos =  4 <  4
		{code: []int{1007, 4, 5, 0, 1, 42}, expected: 1},  // pos < im  =  1 <  5
		{code: []int{1007, 4, 5, 0, 42, 42}, expected: 0}, // pos < im  = 42 <  5
		{code: []int{1107, 4, 5, 0, 1, 42}, expected: 1},  // im  < im  =  4 <  5
		{code: []int{1107, 4, 3, 0, 1, 42}, expected: 0},  // im  < im  =  4 <  3
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpMul(t *testing.T) {
	tt := []struct {
		code     []int
		expected int
	}{
		{code: []int{2, 5, 6, 0, 99, 1, 42}, expected: 42},     // pos * pos = 1 * 42
		{code: []int{102, 5, 6, 0, 99, 1, 42}, expected: 210},  // im * pos = 5 * 42
		{code: []int{1002, 5, 6, 0, 99, 1, 42}, expected: 6},   // pos * im = 1 * 6
		{code: []int{1102, 5, 6, 0, 99, 1, 42}, expected: 30},  // im * im = 5 * 6
		{code: []int{10002, 5, 6, 0, 99, 1, 42}, expected: 42}, // pos * pos = 1 * 42
	}
	cpu := &CPU{}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, 4)
			common.Assert(t, cpu.memory[0] == tc.expected, "Test failed with unexpected result %v, excpected %v", cpu.memory[0], tc.expected)
		})
	}
}

func TestOpOut(t *testing.T) {
	tt := []struct {
		code     []int
		expected int
	}{
		{code: []int{4, 0}, expected: 4},         // pos -> out
		{code: []int{104, 0}, expected: 0},       // im -> out
		{code: []int{1004, 0}, expected: 1004},   // pos -> out
		{code: []int{1104, 0}, expected: 0},      // im -> out
		{code: []int{10004, 0}, expected: 10004}, // pos -> out
	}
	out := -1
	cpu := &CPU{memory: []int{3, 0}, stdOut: func(data int) { out = data }}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test_%v", i), func(t *testing.T) {
			cpu.memory = tc.code
			cpu.pc = 0
			executeAndCheckResult(t, cpu, 2)
			common.Assert(t, out == tc.expected, "Test failed with unexpected result %v, excpected %v", out, tc.expected)
		})
	}
}

func TestOpExit(t *testing.T) {
	//testOp(t, opExit, []int{1, 2, 3}, false, false, 0, 0)
}

/*
type testMemoryAccessor struct {
	staAddr   int
	staValue  int
	calledLda bool
	calledSta bool
}

func (m *testMemoryAccessor) lda(addr int) int {
	m.calledLda = true
	return 42
}
func (m *testMemoryAccessor) sta(addr, value int) {
	m.calledSta = true
	m.staAddr = addr
	m.staValue = value
}

type testIO struct {
	in  int
	out int
}

func (io *testIO) Read() int {
	return io.in
}

func (io *testIO) Write(value int) {
	io.out = value
}


*/
