package intcodecpu

import (
	"fmt"
)

// CPU provieds the basic contents od the CPU/interpreter where a program is executed.
type CPU struct {
	program     []int        // Original program
	memory      []int        // Running program
	pc          uint         // Program Counter
	stdIn       StdIn        // Standard input function
	stdOut      StdOut       // Standard output function
	relBase     int          // Relative base offset for mode 2 OpCode parameters
	data        map[int]int  // Used when program is using addresses outside it's own program space (assumes no code will be executed from here)
	haltProgram bool         // Is set to true when CPU should abort the running program
	DebugOutput func(string) // Set to a function to receive debugging output from CPU processing
}

func (c *CPU) debug(message string) {
	if c.DebugOutput == nil {
		return
	}
	c.DebugOutput(message)
}

func (c *CPU) debugf(format string, v ...interface{}) {
	if c.DebugOutput == nil {
		return
	}
	c.DebugOutput(fmt.Sprintf(format, v...))
}

func (c *CPU) getOpCode() opCode {
	o := opCode(c.memory[c.pc] % 100)
	if _, ok := opCodeStrings[o]; ok {
		return o
	}
	return opERR
}

func (c *CPU) getOpParams(opLen uint) (p1Mode, p1Reg, p2Mode, p2Reg, p3Mode, p3Reg int) {
	p := c.memory[c.pc : c.pc+opLen]
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

func (c *CPU) lda(addr int, mode int) int {
	switch mode {
	case 1:
		return addr
	case 2:
		addr += c.relBase
	}

	if addr < len(c.memory) {
		return c.memory[addr]
	}
	if v, ok := c.data[addr]; ok {
		return v
	}
	c.debug("<?>")
	return 0
}

func (c *CPU) sta(addr, value, mode int) {
	if mode == 2 {
		addr += c.relBase
	}
	if addr < len(c.memory) {
		c.memory[addr] = value
		c.debugf(" = %v => %v\n", value, addr)
	} else {
		c.data[addr] = value
		c.debugf(" = %v => %v[*]\n", value, addr)
	}
}

// DumpMemory return a dump of the active memory, starting from the start address
func (c *CPU) DumpMemory(startAddress, length int) (dump []int) {
	dump = make([]int, length)
	for i := range dump {
		dump[i] = c.lda(startAddress+i, 0)
	}
	return
}

// Halt is called to stop the running program
func (c *CPU) Halt() {
	c.haltProgram = true
}

// PatchMemory is used to change the memory, the patched memory is reverted when Reset() is called
func (c *CPU) PatchMemory(startAddress int, patchData []int) {
	for i, d := range patchData {
		c.sta(startAddress+i, d, 0)
	}
}

// Reset will reset both program and pc to it's original states
func (c *CPU) Reset() {
	copy(c.memory, c.program)
	c.pc = 0
	c.relBase = 0
	c.data = make(map[int]int)
}

// Run start the program
func (c *CPU) Run() (err error) {
	c.debug("--program run--\n")
	c.haltProgram = false
	done := false
	for !done {
		if done, err = c.executeOpCode(); err != nil {
			c.debug("--program fault--\n")
			return
		}
		if c.haltProgram {
			done = true
		}
	}
	c.debug("--program end--\n")
	return
}
