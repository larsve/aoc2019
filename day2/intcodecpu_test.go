package main

import (
	"testing"

	"aoc2019/common"
)

func TestGetOpCode(t *testing.T) {
	i := sampleProgram()

	testOpCode := func(actual, expected opCode) {
		common.Assert(t, actual == expected, "Not the expected opCode, expected %v but got %v", expected, actual)
		i.pc += 4
	}

	testOpCode(i.getOpCode(), opAdd)
	testOpCode(i.getOpCode(), opMul)
	testOpCode(i.getOpCode(), opExit)
}

func TestProcessNextOpCode(t *testing.T) {
	p := sampleProgram()

	test := func(expectedDone, expectedError bool) {
		d, e := p.processNextOpCode()
		common.Assert(t, d == expectedDone, "Not the expected done, expected %v but got %v", expectedDone, d)
		common.Assert(t, (e != nil) == expectedError, "Not the expected error, expected %v but got %v", expectedError, e)
	}

	test(false, false)
	test(false, false)
	test(true, false)

	p = faultyProgram()
	test(false, true)
}

func TestReset(t *testing.T) {
	p := faultyProgram()
	p.program[2] = 42
	p.reset()
	common.Assert(t, p.program[2] == 0, "Reset failed")
}

func TestRun(t *testing.T) {
	p := smallAddProgram()
	r,e := p.run()
	common.Assert(t, e == nil, "Run ended with an error: %v", e)
	common.Assert(t, r == 2, "Run didn't give the expected result, expected 2 but got %v", r)

	p = faultyProgram()
	r,e = p.run()
	common.Assert(t, e != nil, "Run didn't end with an error")
}

func TestSetAlarmCode(t *testing.T) {
	i := sampleProgram()
	test := func(code, expectedNoun, expectedVerb int) {
		i.setAlarmCode(code)
		common.Assert(t, i.program[1] == expectedNoun, "Not the expected noun, expected %v but got %v (alarm code: %v)", expectedNoun, i.program[1], code)
		common.Assert(t, i.program[2] == expectedVerb, "Not the expected verb, expected %v but got %v (alarm code: %v)", expectedVerb, i.program[2], code)
	}
	test(99, 00, 99)
	test(9900, 99, 00)
	test(999, 9, 99)
	test(1000, 10, 0)
}

func TestGetNounAndVerbFor(t *testing.T) {
	common.Assert(t, getNounAndVerbFor(19690720) == 6533, "GetNounAndVerbFor falied")
}

func BenchmarkGetNounAndVerbFor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getNounAndVerbFor(19690720)
	}
}

func TestCpuGetNounAndVerbFor(t *testing.T) {
	c := newProgram(code[:])
	common.Assert(t, c.getNounAndVerbFor(19690720) == 6533, "GetNounAndVerbFor falied")
}
func BenchmarkCpuGetNounAndVerbFor(b *testing.B) {
	c := newProgram(code[:])
	for n := 0; n < b.N; n++ {
		c.getNounAndVerbFor(19690720)
	}
}

func sampleProgram() *cpu {
	return newProgram([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
}
func smallAddProgram() *cpu {
	return newProgram([]int{1, 0, 0, 0, 99})
}
func faultyProgram() *cpu {
	return newProgram([]int{0, 0, 0, 0, 42})
}
