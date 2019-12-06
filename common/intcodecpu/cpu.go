package intcodecpu

import (
	"fmt"
)

// CPU provieds the basic contents od the CPU/interpreter where a program is executed.
type CPU struct {
	program []int
	memory  []int
	pc      uint
	stdIn   StdIn
	stdOut  StdOut
}

func (c *CPU) lda(addr int) int {
	return c.memory[addr]
}

func (c *CPU) sta(addr, value int) {
	c.memory[addr] = value
}

func (c *CPU) reset() {
	copy(c.memory, c.program)
	c.pc = 0
}

// Run start the program
func (c *CPU) Run() (result int, err error) {
	done := false
	result = 0
	for !done {
		if done, err = c.executeOpCode(); err != nil {
			return
		}
	}
	return c.memory[0], nil
}

// SetAlarmCode will set the 2 first parameter of the first operand in the program
func (c *CPU) SetAlarmCode(code int) {
	c.memory[1] = code / 100
	c.memory[2] = code % 100
}

// GetNounAndVerbFor return the "Noun" and "Verb" (Alarm code) for a specified program output
func (c *CPU) GetNounAndVerbFor(output int) int {
	for i := 0; i <= 9999; i++ {
		c.reset()
		c.SetAlarmCode(i)
		r, e := c.Run()
		if e != nil {
			fmt.Printf("Alarm Code %v broke interpreter (%v), oh well, checking next..\n", i, e)
		}
		if r == output {
			return i
		}
	}
	return -1
}
