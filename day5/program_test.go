package main

import (
	"reflect"
	"testing"

	"aoc2019/common"
	cpu "aoc2019/common/intcodecpu"
)

func TestGetDiagnosticCodePart1(t *testing.T) {
	expected := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 9219874}
	c, e := getDiagnosticCode(1)
	common.Assert(t, e == nil, "Diagnostic failed with error %v", e)
	common.Assert(t, len(c) == 10, "Diagnostic failed unexpected output %v", c)
	common.Assert(t, reflect.DeepEqual(c, expected), "Diagnostic failed, expected %v, got %v", expected, c)
}

func TestGetDiagnosticCodePart2(t *testing.T) {
	expected := 5893654
	c, e := getDiagnosticCode(5)
	common.Assert(t, e == nil, "Diagnostic failed with error %v", e)
	common.Assert(t, len(c) == 1, "Diagnostic failed unexpected output %v", c)
	common.Assert(t, c[0] == expected, "Diagnostic failed, expected %v, got %v", expected, c)
}

func BenchmarkGetDiagnosticCodePart2(b *testing.B) {
	c := cpu.NewProgramWithIO(diagnosticProgram, func() int { return 5 }, func(data int) {})
	for n := 0; n < b.N; n++ {
		e := c.Run()
		common.Assert(b, e == nil, "TestDay5Part2 falied with error: %v", e)
		c.Reset()
	}
}
