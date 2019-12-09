package main

import (
	"testing"

	"aoc2019/common"
)

var (
	exampleProgram1 = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	exampleProgram2 = []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	exampleProgram3 = []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
)

func TestCalcThrusterSignal(t *testing.T) {
	tt := []struct {
		name          string // Name of sub test
		program       []int  // IntCode program to run
		phase         int    // Phase setting to test
		trusterSignal int    // Expected thruster signal
	}{
		{name: "Example1", program: exampleProgram1, phase: 43210, trusterSignal: 43210},
		{name: "Example2", program: exampleProgram2, phase: 1234, trusterSignal: 54321},
		{name: "Example3", program: exampleProgram3, phase: 10432, trusterSignal: 65210},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ac := newAmpController(tc.program)
			ts := ac.calcThrusterSignal(tc.phase)
			common.Assert(t, ts == tc.trusterSignal, "Thruster signal mismatch for %v, got %v, but expected %v", tc.name, ts, tc.trusterSignal)
		})
	}
}

func TestGeneratePossiblePhaseSettings(t *testing.T) {
	ps := generatePossiblePhaseSettings()
	common.Assert(t, len(ps) == 120, "GeneratePossiblePhaseSettings failed, got %v combinations, but expected 120", len(ps))
}

func TestGetLargestOutputSignal(t *testing.T) {
	sig := getLargestOutputSignalPart1(ampProgram)
	common.Assert(t, sig == 567045, "Not the expected truster signal, got %v, expected 567045", sig)
}
