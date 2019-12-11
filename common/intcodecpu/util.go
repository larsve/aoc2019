package intcodecpu

// NewProgram creates a new CPU/program interpreter loaded with provided program.
func NewProgram(code []int) *CPU {
	var memory = make([]int, len(code))
	copy(memory, code)
	return &CPU{program: code, memory: memory[:], pc: 0, data: make(map[int]int)}
}

// NewProgramWithIO creates a new CPU/program interpreter loaded with provided program and I/O.
func NewProgramWithIO(code []int, in StdIn, out StdOut) *CPU {
	var memory = make([]int, len(code))
	copy(memory, code)
	return &CPU{program: code, memory: memory[:], pc: 0, stdIn: in, stdOut: out, data: make(map[int]int)}
}
