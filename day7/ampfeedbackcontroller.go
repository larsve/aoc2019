package main

import (
	"fmt"
	"sync"

	"aoc2019/common/intcodecpu"
)

type ampFeedbackController struct {
	cpu         [5]*intcodecpu.CPU
	stdIn       [5]chan int
	ampInOutput int
}

func (a *ampFeedbackController) calcThrusterSignal(phaseSetting int) int {
	// Load each amp input with corresponding digit from phase setting..
	a.stdIn[0] <- phaseSetting / 10000
	a.stdIn[1] <- phaseSetting / 1000 % 10
	a.stdIn[2] <- phaseSetting / 100 % 10
	a.stdIn[3] <- phaseSetting / 10 % 10
	a.stdIn[4] <- phaseSetting % 10

	// Load the first amp signal
	a.stdIn[0] <- 0

	// Start amplifiers
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(ampNo int) {
			defer wg.Done()
			a.cpu[ampNo].Reset()
			if err := a.cpu[ampNo].Run(); err != nil {
				fmt.Printf("AMP#%v abotred due to %v", ampNo, err)
			}
		}(i)
	}
	wg.Wait()

	// Read last AMP output, from AMP #0 StdIn
	a.ampInOutput = <-a.stdIn[0]

	//fmt.Printf("Phase setting %v = %v\n", phaseSetting, a.ampInOutput)
	return a.ampInOutput
}

func newAmpFeedbackController(program []int) *ampFeedbackController {
	a := &ampFeedbackController{stdIn: [5]chan int{make(chan int, 2), make(chan int, 1), make(chan int, 1), make(chan int, 1), make(chan int, 1)}}
	createStdInFunc := func(stdIn chan int) func() int { return func() int { return <-stdIn } }
	createStdOutFunc := func(stdOut chan int) func(int) { return func(data int) { stdOut <- data } }
	for i := 0; i < 5; i++ {
		a.cpu[i] = intcodecpu.NewProgramWithIO(program, createStdInFunc(a.stdIn[i]), createStdOutFunc(a.stdIn[(i+1)%5]))
	}
	return a
}

func getLargestOutputSignalPart2(program []int) int {
	return getLargestOutputSignal(program, func(psChan chan int) {
		for _, ps := range generatePossiblePhaseSettings() {
			psChan <- ps + 55555 // Ugly code re-use by "Converting" phase setting 01234 -> 56789 for the second part...
		}
	}, func(program []int) ampController { return newAmpFeedbackController(program) })
}
