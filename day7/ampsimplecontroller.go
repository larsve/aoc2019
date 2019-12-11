package main

import (
	"fmt"

	"aoc2019/common/intcodecpu"
)

type ampSimpleController struct {
	cpu          *intcodecpu.CPU
	phaseSetting [5]int
	ampIteration int
	inputCnt     int
	ampInOutput  int
}

func (a *ampSimpleController) stdIn() int {
	a.inputCnt++
	if a.inputCnt == 1 {
		return a.phaseSetting[a.ampIteration]
	}
	return a.ampInOutput
}

func (a *ampSimpleController) stdOut(data int) {
	a.ampInOutput = data
}

func (a *ampSimpleController) calcThrusterSignal(phaseSetting int) int {
	a.phaseSetting[0] = phaseSetting / 10000
	a.phaseSetting[1] = phaseSetting / 1000 % 10
	a.phaseSetting[2] = phaseSetting / 100 % 10
	a.phaseSetting[3] = phaseSetting / 10 % 10
	a.phaseSetting[4] = phaseSetting % 10
	a.ampInOutput = 0
	for a.ampIteration = 0; a.ampIteration < 5; a.ampIteration++ {
		a.inputCnt = 0
		a.cpu.Reset()
		if err := a.cpu.Run(); err != nil {
			fmt.Printf("Amp iteration %v aborted du to %v\n", a.ampIteration, err)
			return -1
		}
	}
	return a.ampInOutput
}

func newAmpController(program []int) *ampSimpleController {
	a := &ampSimpleController{}
	a.cpu = intcodecpu.NewProgramWithIO(program, a.stdIn, a.stdOut)
	return a
}

func getLargestOutputSignalPart1(program []int) int {
	return getLargestOutputSignal(program, func(psChan chan int) {
		for _, ps := range generatePossiblePhaseSettings() {
			psChan <- ps
		}
	}, func(program []int) ampController { return newAmpController(program) })
}
