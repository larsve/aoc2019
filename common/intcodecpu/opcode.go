package intcodecpu

import "fmt"

type opCode byte

const (
	opErr opCode = iota
	opAdd
	opMul
	opIn
	opOut
	opJmpT
	opJmpF
	opLess
	opEqu
	opExit opCode = 99
)

var opCodeStrings = map[opCode]string{
	opErr:  "ERR",
	opAdd:  "ADD",
	opMul:  "MUL",
	opIn:   "IN",
	opOut:  "OUT",
	opJmpT: "JMPT",
	opJmpF: "JMPF",
	opLess: "LESS",
	opEqu:  "EQU",
	opExit: "HALT",
}

func (o opCode) String() string {
	if s, ok := opCodeStrings[o]; ok {
		return s
	}
	return "Unknown OpCode"
}

func (c *CPU) getOpCode() opCode {
	o := opCode(c.memory[c.pc] % 100)
	if _, ok := opCodeStrings[o]; ok {
		return o
	}
	return opErr
}

func (c *CPU) opAdd() {
	p1m, p1, p2m, p2, _, p3 := decodeOpParams(c.memory[c.pc : c.pc+4])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	if p2m == 0 {
		p2 = c.lda(p2)
	}
	c.sta(p3, p1+p2)
	c.pc += 4
}

func (c *CPU) opEqu() {
	p1m, p1, p2m, p2, _, p3 := decodeOpParams(c.memory[c.pc : c.pc+4])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	if p2m == 0 {
		p2 = c.lda(p2)
	}
	if p1 == p2 {
		c.sta(p3, 1)
	} else {
		c.sta(p3, 0)
	}
	c.pc += 4
}

func (c *CPU) opIn() {
	p1 := c.memory[c.pc+1] // Always position mode
	c.sta(p1, c.stdIn())
	c.pc += 2
}

func (c *CPU) opJmpF() {
	p1m, p1, p2m, p2, _, _ := decodeOpParams(c.memory[c.pc : c.pc+3])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	if p1 == 0 {
		if p2m == 0 {
			p2 = c.lda(p2)
		}
		c.pc = uint(p2)
	} else {
		c.pc += 3
	}
}

func (c *CPU) opJmpT() {
	p1m, p1, p2m, p2, _, _ := decodeOpParams(c.memory[c.pc : c.pc+3])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	if p1 != 0 {
		if p2m == 0 {
			p2 = c.lda(p2)
		}
		c.pc = uint(p2)
	} else {
		c.pc += 3
	}
}

func (c *CPU) opLess() {
	p1m, p1, p2m, p2, _, p3 := decodeOpParams(c.memory[c.pc : c.pc+4])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	if p2m == 0 {
		p2 = c.lda(p2)
	}
	if p1 < p2 {
		c.sta(p3, 1)
	} else {
		c.sta(p3, 0)
	}
	c.pc += 4
}

func (c *CPU) opMul() {
	p1m, p1, p2m, p2, _, p3 := decodeOpParams(c.memory[c.pc : c.pc+4])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	if p2m == 0 {
		p2 = c.lda(p2)
	}
	c.sta(p3, p1*p2)
	c.pc += 4
}

func (c *CPU) opOut() {
	p1m, p1, _, _, _, _ := decodeOpParams(c.memory[c.pc : c.pc+2])
	if p1m == 0 {
		p1 = c.lda(p1)
	}
	c.stdOut(p1)
	c.pc += 2
}

func (c *CPU) executeOpCode() (done bool, err error) {
	opCode := c.getOpCode()
	switch opCode {
	case opAdd:
		c.opAdd()
		return false, nil
	case opEqu:
		c.opEqu()
		return false, nil
	case opIn:
		c.opIn()
		return false, nil
	case opJmpF:
		c.opJmpF()
		return false, nil
	case opJmpT:
		c.opJmpT()
		return false, nil
	case opLess:
		c.opLess()
		return false, nil
	case opMul:
		c.opMul()
		return false, nil
	case opOut:
		c.opOut()
		return false, nil
	case opExit:
		return true, nil
	}
	return true, fmt.Errorf("Unknown/unhandled OpCode %v (%v) at address %v", c.memory[c.pc], opCode, c.pc)
}
