package main

import (
	"errors"
	"fmt"
)

type cpu struct {
	originalProgram []int
	program         []int
	pc              uint
	debug           bool
}

var code = [...]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 2, 13, 23, 27, 2, 27, 13, 31, 2, 31, 10, 35, 1, 6, 35, 39, 1, 5, 39, 43, 1, 10, 43, 47, 1, 5, 47, 51, 1, 13, 51, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 13, 63, 67, 1, 6, 67, 71, 1, 71, 10, 75, 2, 13, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 6, 87, 91, 1, 91, 13, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 5, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 2, 115, 119, 1, 119, 6, 0, 99, 2, 0, 14, 0}

func (c *cpu) getOpCode() opCode {
	switch c.program[c.pc] {
	case 1:
		return opAdd
	case 2:
		return opMul
	case 99:
		return opExit
	}
	return opErr
}
func (c *cpu) debugf(msg string, v ...interface{}) {
	if c.debug {
		fmt.Printf(msg, v...)
	}
}
func (c *cpu) debugln(msg string) {
	if c.debug {
		fmt.Println(msg)
	}
}
func (c *cpu) lda(addr int) int {
	return c.program[addr]
}
func (c *cpu) sta(addr, value int) {
	c.program[addr] = value
}
func (c *cpu) processNextOpCode() (done bool, err error) {
	done = false
	opCode := c.getOpCode()
	c.debugf("PC: %v\tOP: %v", c.pc, opCode)
	if opCode == opExit {
		c.debugln("")
		return true, nil
	}
	if opCode == opErr {
		c.debugln("")
		return false, errors.New("Unknown opcode detected")
	}
	opParams := c.program[c.pc+1 : c.pc+4]
	c.debugf("(%v)\n", opParams)
	err = opCode.execute(opParams, c)
	c.pc += 4
	return
}

func (c *cpu) reset() {
	copy(c.program, c.originalProgram)
	c.pc = 0
}

func (c *cpu) run() (result int, err error) {
	done := false
	result = 0
	for !done {
		if done, err = c.processNextOpCode(); err != nil {
			return
		}
	}
	return c.program[0], nil
}

func (c *cpu) setAlarmCode(code int) {
	c.program[1] = code / 100
	c.program[2] = code % 100
}

func newProgram(code []int) *cpu {
	var memory = make([]int, len(code))
	copy(memory, code)
	return &cpu{originalProgram: code, program: memory[:], pc: 0}
}

func getNounAndVerbFor(output int) int {
	p := newProgram(code[:])
	for i := 0; i <= 9999; i++ {
		p.setAlarmCode(i)
		r, e := p.run()
		if e != nil {
			fmt.Printf("Alarm Code %v broke interpreter (%v), oh well, checking next..\n", i, e)
		}
		if r == output {
			return i
		}
		p.reset()
	}
	return -1
}

func (c *cpu) getNounAndVerbFor(output int) int {
	for i := 0; i <= 9999; i++ {
		c.reset()
		c.setAlarmCode(i)
		r, e := c.run()
		if e != nil {
			fmt.Printf("Alarm Code %v broke interpreter (%v), oh well, checking next..\n", i, e)
		}
		if r == output {
			return i
		}
	}
	return -1
}
