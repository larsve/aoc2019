package main

import (
	"testing"

	"aoc2019/common"
)

var (
	exampleProgram4 = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	exampleProgram5 = []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}
)

func TestCalcThrusterSignalPart2(t *testing.T) {
	tt := []struct {
		name          string // Name of sub test
		program       []int  // IntCode program to run
		phase         int    // Phase setting to test
		trusterSignal int    // Expected thruster signal
	}{
		// Re-test Example program 1-3 with the feedback controller, as well as the new (part2) programs
		{name: "Example1", program: exampleProgram1, phase: 43210, trusterSignal: 43210},
		{name: "Example2", program: exampleProgram2, phase: 1234, trusterSignal: 54321},
		{name: "Example3", program: exampleProgram3, phase: 10432, trusterSignal: 65210},
		{name: "Example4", program: exampleProgram4, phase: 98765, trusterSignal: 139629729},
		{name: "Example5", program: exampleProgram5, phase: 97856, trusterSignal: 18216},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ac := newAmpFeedbackController(tc.program)
			ts := ac.calcThrusterSignal(tc.phase)
			common.Assert(t, ts == tc.trusterSignal, "Thruster signal mismatch for %v, got %v, but expected %v", tc.name, ts, tc.trusterSignal)
		})
	}
}

func TestGetLargestOutputSignalPart2(t *testing.T) {
	sig := getLargestOutputSignalPart2(ampProgram)
	common.Assert(t, sig == 39016654, "Not the expected truster signal, got %v, expected 39016654", sig)
}
