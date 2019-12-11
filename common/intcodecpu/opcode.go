package intcodecpu

import "fmt"

type opCode byte

const (
	opERR opCode = iota // Op code error
	opADD               // Add
	opMUL               // Multiplicate
	opINP               // Input
	opOUT               // Output
	opJMT               // Jump if True
	opJMF               // Jump if Flase
	opLES               // Less than
	opEQU               // Equal to
	opCRB               // Change Relative Base offset
	opHLT opCode = 99   // Program halt (exit program)
)

const (
	oneParamwOutln  = "(%v) = %v => %v\n"
	twoParams       = "(%v, %v)"
	twoParamsln     = "(%v, %v)\n"
	twoParamswOutln = "(%v, %v) = %v => %v\n"
)

var opCodeStrings = map[opCode]string{
	opERR: "ERR",
	opADD: "ADD",
	opMUL: "MUL",
	opINP: "INP",
	opOUT: "OUT",
	opJMT: "JMT",
	opJMF: "JMF",
	opLES: "LES",
	opEQU: "EQU",
	opCRB: "CRB",
	opHLT: "HLT",
}

func (o opCode) String() string {
	if s, ok := opCodeStrings[o]; ok {
		return s
	}
	return "Unknown OpCode"
}

func (c *CPU) opADD() (bool, error) {
	opLen := uint(4)
	p1m, p1, p2m, p2, p3m, p3 := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	p2 = c.lda(p2, p2m)
	c.debugf(twoParams, p1, p2)
	c.sta(p3, p1+p2, p3m)
	c.pc += opLen
	return false, nil
}

func (c *CPU) opCRB() (bool, error) {
	opLen := uint(2)
	p1m, p1, _, _, _, _ := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	c.debugf(oneParamwOutln, p1, c.relBase+p1, "RB")
	c.relBase += p1
	c.pc += opLen
	return false, nil
}

func (c *CPU) opEQU() (bool, error) {
	opLen := uint(4)
	p1m, p1, p2m, p2, p3m, p3 := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	p2 = c.lda(p2, p2m)
	c.debugf(twoParams, p1, p2)
	if p1 == p2 {
		c.sta(p3, 1, p3m)
	} else {
		c.sta(p3, 0, p3m)
	}
	c.pc += opLen
	return false, nil
}

func (c *CPU) opHLT() (bool, error) {
	c.debug("\n")
	return true, nil
}

func (c *CPU) opINP() (bool, error) {
	opLen := uint(2)
	p1m, p1, _, _, _, _ := c.getOpParams(opLen)
	in := c.stdIn()
	c.debug("()")
	c.sta(p1, in, p1m)
	c.pc += opLen
	return false, nil
}

func (c *CPU) opJMF() (bool, error) {
	opLen := uint(3)
	p1m, p1, p2m, p2, _, _ := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	p2 = c.lda(p2, p2m)
	if p1 == 0 {
		c.pc = uint(p2)
		c.debugf(twoParamswOutln, p1, p2, c.pc, "PC")
	} else {
		c.pc += opLen
		c.debugf(twoParamsln, p1, p2)
	}
	return false, nil
}

func (c *CPU) opJMT() (bool, error) {
	opLen := uint(3)
	p1m, p1, p2m, p2, _, _ := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	p2 = c.lda(p2, p2m)
	if p1 != 0 {
		c.pc = uint(p2)
		c.debugf(twoParamswOutln, p1, p2, c.pc, "PC")
	} else {
		c.pc += opLen
		c.debugf(twoParamsln, p1, p2)
	}
	return false, nil
}

func (c *CPU) opLES() (bool, error) {
	opLen := uint(4)
	p1m, p1, p2m, p2, p3m, p3 := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	p2 = c.lda(p2, p2m)
	c.debugf(twoParams, p1, p2)
	if p1 < p2 {
		c.sta(p3, 1, p3m)
	} else {
		c.sta(p3, 0, p3m)
	}
	c.pc += opLen
	return false, nil
}

func (c *CPU) opMUL() (bool, error) {
	opLen := uint(4)
	p1m, p1, p2m, p2, p3m, p3 := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	p2 = c.lda(p2, p2m)
	c.debugf(twoParams, p1, p2)
	c.sta(p3, p1*p2, p3m)
	c.pc += opLen
	return false, nil
}

func (c *CPU) opOUT() (bool, error) {
	opLen := uint(2)
	p1m, p1, _, _, _, _ := c.getOpParams(opLen)
	p1 = c.lda(p1, p1m)
	c.debugf(oneParamwOutln, p1, p1, "OUT")
	c.stdOut(p1)
	c.pc += opLen
	return false, nil
}

func (c *CPU) executeOpCode() (done bool, err error) {
	opCode := c.getOpCode()
	c.debugf("PC %4v : RB %4v : [%5v]  %v", c.pc, c.relBase, c.memory[c.pc], opCode)
	switch opCode {
	case opADD:
		return c.opADD()
	case opCRB:
		return c.opCRB()
	case opEQU:
		return c.opEQU()
	case opINP:
		return c.opINP()
	case opJMF:
		return c.opJMF()
	case opJMT:
		return c.opJMT()
	case opLES:
		return c.opLES()
	case opMUL:
		return c.opMUL()
	case opOUT:
		return c.opOUT()
	case opHLT:
		return c.opHLT()
	}
	return true, fmt.Errorf("Unknown/unhandled OpCode %v (%v) at address %v", c.memory[c.pc], opCode, c.pc)
}
