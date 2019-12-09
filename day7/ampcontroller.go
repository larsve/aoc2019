package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

var ampProgram = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 51, 76, 93, 110, 191, 272, 353, 434, 99999, 3, 9, 1002, 9, 2, 9, 1001, 9, 3, 9, 1002, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 1001, 9, 4, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 5, 9, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}

type ampController interface {
	calcThrusterSignal(phaseSetting int) int
}

func generatePossiblePhaseSettings() []int {
	// It's only 120 combinations, so hard code it for now...
	return []int{
		1234, 1243, 1324, 1342, 1423, 1432, 2134, 2143,
		2314, 2341, 2413, 2431, 3124, 3142, 3214, 3241,
		3412, 3421, 4123, 4132, 4213, 4231, 4312, 4321,
		10234, 10243, 10324, 10342, 10423, 10432, 12034, 12043,
		12304, 12340, 12403, 12430, 13024, 13042, 13204, 13240,
		13402, 13420, 14023, 14032, 14203, 14230, 14302, 14320,
		20134, 20143, 20314, 20341, 20413, 20431, 21034, 21043,
		21304, 21340, 21403, 21430, 23014, 23041, 23104, 23140,
		23401, 23410, 24013, 24031, 24103, 24130, 24301, 24310,
		30124, 30142, 30214, 30241, 30412, 30421, 31024, 31042,
		31204, 31240, 31402, 31420, 32014, 32041, 32104, 32140,
		32401, 32410, 34012, 34021, 34102, 34120, 34201, 34210,
		40123, 40132, 40213, 40231, 40312, 40321, 41023, 41032,
		41203, 41230, 41302, 41320, 42013, 42031, 42103, 42130,
		42301, 42310, 43012, 43021, 43102, 43120, 43201, 43210,
	}
}

func getLargestOutputSignal(program []int, sendPhaseSettingsFunc func(chan int), newControllerFunc func([]int) ampController) int {
	startTime := time.Now()
	maxSignal := math.MinInt32
	psChan := make(chan int, 25) // Phase setting
	resChan := make(chan [2]int) // Phase setting and thruster signal
	wg := sync.WaitGroup{}       // Used to close resChan when last phase setting have been processed by a worker

	// Start adding phase settings to channel..
	go func() {
		fmt.Println("Start sending phase settings..")
		sendPhaseSettingsFunc(psChan)
		close(psChan) // Close channel to indicate that all combinations have been sent..
		fmt.Println("All phase settings sent.")
	}()

	// Start worker go-routines that will process phase settings from psChan
	nproc := runtime.NumCPU()
	wg.Add(nproc)
	for i := 0; i < nproc; i++ {
		go func(wno int) {
			fmt.Printf("Worker#%v starting..\n", wno)
			defer wg.Done()
			a := newControllerFunc(program)
			for ps := range psChan {
				resChan <- [2]int{ps, a.calcThrusterSignal(ps)}
			}
			fmt.Printf("Worker#%v stopped.\n", wno)
		}(i)
	}

	// Start a go-routine that watches wg and closes resChan when all phase settings have been processed by a worker
	go func() {
		fmt.Println("Waiting for workers to stop..")
		wg.Wait()
		fmt.Println("All workers stopped, closing result channel..")
		close(resChan)
	}()

	// Collect all answers as they start to come back
	fmt.Println("Start collecting thruster signals from workers..")
	for res := range resChan {
		if res[1] > maxSignal {
			maxSignal = res[1]
		}
	}
	computeTime := time.Since(startTime)
	fmt.Printf("All thruster signals received.. Max thruster signal: %v, Compute time: %v\n", maxSignal, computeTime)
	return maxSignal
}
