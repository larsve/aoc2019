package intcodecpu

func decodeOpParams(p []int) (p1Mode, p1Reg, p2Mode, p2Reg, p3Mode, p3Reg int) {
	p1Mode = p[0] / 100 % 10
	p2Mode = p[0] / 1000 % 10
	p3Mode = p[0] / 10000 % 10
	pSize := len(p)
	if pSize > 1 {
		p1Reg = p[1]
	}
	if pSize > 2 {
		p2Reg = p[2]
	}
	if pSize > 3 {
		p3Reg = p[3]
	}
	return
}

// NewProgram creates a new CPU/program interpreter loaded with provided program.
func NewProgram(code []int) *CPU {
	var memory = make([]int, len(code))
	copy(memory, code)
	return &CPU{program: code, memory: memory[:], pc: 0}
}

// NewProgramWithIO creates a new CPU/program interpreter loaded with provided program and I/O.
func NewProgramWithIO(code []int, in StdIn, out StdOut) *CPU {
	var memory = make([]int, len(code))
	copy(memory, code)
	return &CPU{program: code, memory: memory[:], pc: 0, stdIn: in, stdOut: out}
}
