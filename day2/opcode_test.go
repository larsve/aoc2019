package main

import (
	"testing"

	"aoc2019/common"
)

func TestOpCodeToString(t *testing.T) {
	common.Assert(t, opAdd.String() == "ADD", "opAdd did not return expected string")
	common.Assert(t, opMul.String() == "MUL", "opMul did not return expected string")
	common.Assert(t, opErr.String() == "ERR", "opErr did not return expected string")
	common.Assert(t, opExit.String() == "BRK", "opExit did not return expected string")
	common.Assert(t, opCode(42).String() == "Unknown OpCode", "Invalid opCode did not return expected string")
}

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

func testOp(t *testing.T, o opCode, p []int, elda, esta bool, eaddr, eval int) {
	m := &testMemoryAccessor{}
	e := o.execute(p, m)
	common.Assert(t, e == nil, "%v terminated with an error: %v", o, e)
	common.Assert(t, m.calledLda == elda, "%v LDA call, expected %v", o, elda)
	common.Assert(t, m.calledSta == esta, "%v STA call, expected %v", o, esta)
	common.Assert(t, m.staAddr == eaddr, "%v STA addr, expected %v, but were %v", o, eaddr, m.staAddr)
	common.Assert(t, m.staValue == eval, "%v STA value, expected %v, but were %v", o, eaddr, m.staValue)
}
func TestOpErrExecute(t *testing.T) {
	testOp(t, opErr, []int{1, 2, 3}, false, false, 0, 0)
}

func TestOpAddExecute(t *testing.T) {
	testOp(t, opAdd, []int{1, 2, 3}, true, true, 3, 84)
}

func TestOpMulExecute(t *testing.T) {
	testOp(t, opMul, []int{1, 2, 42}, true, true, 42, 1764)
}

func TestOpExitExecute(t *testing.T) {
	testOp(t, opExit, []int{1, 2, 3}, false, false, 0, 0)
}
