package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

type signal []byte

const (
	inputSignal = "59766832516471105169175836985633322599038555617788874561522148661927081324685821180654682056538815716097295567894852186929107230155154324411726945819817338647442140954601202408433492208282774032110720183977662097053534778395687521636381457489415906710702497357756337246719713103659349031567298436163261681422438462663511427616685223080744010014937551976673341714897682634253850270219462445161703240957568807600494579282412972591613629025720312652350445062631757413159623885481128914333982571503540357043736821931054029305931122179293220911720263006705242490442826574028623201238659548887822088996956559517179003476743001815465428992906356931239533104"
)

var (
	pattern = []int{0, 1, 0, -1}
)

func (s signal) String() string {
	r := make([]byte, len(s))
	for i, b := range s {
		r[i] = b + 48
	}
	return string(r)
}

func (s signal) Value() string {
	l := min(len(s), 8)
	r := make([]byte, l)
	for i := 0; i < l; i++ {
		r[i] = s[i] + 48
	}
	return string(r)
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// stos convert a string to a signal
func stos(input string) signal {
	res := make(signal, len(input))
	for i, r := range input {
		res[i] = (byte(r) - 48) % 10
	}
	return res
}

type iterationResult struct {
	i int
	r byte
}

func fft(input signal) signal {
	out := make(signal, len(input))
	work := make(chan int, 1)
	irc := make(chan iterationResult)

	// Put work on the work queue
	go func() {
		for i := range input {
			work <- i
		}
		close(work)
	}()

	// Create workers
	wg := &sync.WaitGroup{}
	for i := 0; i < min(runtime.NumCPU(), len(input)); i++ {
		wg.Add(1)
		go func(wid int) {
			for iteration := range work {
				var bo int
				for j := iteration; j <= len(input)-1; j++ {
					bi := int(input[j])
					pi := ((j + 1) / (iteration + 1)) % 4
					bo += bi * pattern[pi]
				}
				irc <- iterationResult{i: iteration, r: byte(abs(bo) % 10)}
			}
			wg.Done()
		}(i)
	}

	// Wait for all workers to complete all work
	go func() {
		wg.Wait()
		close(irc)
	}()

	// Collect results from workers
	for itRes := range irc {
		out[itRes.i] = itRes.r
	}
	return out
}

func fftLoop(input signal, phases int) signal {
	signal := input
	for j := 0; j < phases; j++ {
		signal = fft(signal)
	}
	return signal
}

func part1() signal {
	start := time.Now()
	sig := fftLoop(stos(inputSignal), 100)
	et := time.Since(start)
	fmt.Printf("output signal: %s, execution time: %v\n", sig.Value(), et)
	return sig
}

func part2() signal {
	start := time.Now()
	sig := fftLoop(stos(strings.Repeat(inputSignal, 10000)), 100)
	et := time.Since(start)
	fmt.Printf("output signal: %s, execution time: %v\n", sig.Value(), et)
	return sig
}
