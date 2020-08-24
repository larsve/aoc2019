package main

import (
	"aoc2019/common/intcodecpu"
	"fmt"
)

type (
	tile int
	pos  struct {
		x, y int
	}
	arcadeMachine struct {
		cpu    intcodecpu.CPU
		ball   pos
		paddle pos
		pCnt   int
		xPos   int
		yPos   int
		score  int
		screen map[pos]tile
	}
)

const (
	empty tile = iota
	wall
	block
	hPaddle
	ball
)

func (t tile) String() string {
	return [...]string{"Empty", "Wall", "Block", "Horizontal Paddle", "Ball"}[t]
}

func (a *arcadeMachine) cpuInput() int {
	// Just follow the ball...
	switch ofs := a.ball.x - a.paddle.x; {
	case ofs > 0:
		return 1
	case ofs < 0:
		return -1
	}
	return 0
}

func (a *arcadeMachine) cpuOutput(data int) {
	switch a.pCnt {
	case 0:
		a.xPos = data
		a.pCnt++
	case 1:
		a.yPos = data
		a.pCnt++
	case 2:
		a.pCnt = 0
		if a.xPos == -1 && a.yPos == 0 {
			a.score = data
		} else {
			p := pos{x: a.xPos, y: a.yPos}
			t := tile(data)
			a.screen[p] = t
			/*
				if t == ball || t == hPaddle {
					fmt.Printf("%v @ %v\n", t, p)
				}
			*/
			if t == ball {
				a.ball = p
			} else if t == hPaddle {
				a.paddle = p
			}
		}
	}
}

func (a *arcadeMachine) run() error {
	return a.cpu.Run()
}

func newArcadeMachie() *arcadeMachine {
	r := &arcadeMachine{ball: pos{0, 0}, paddle: pos{0, 0}, screen: make(map[pos]tile)}
	r.cpu = *intcodecpu.NewProgramWithIO(intCode, r.cpuInput, r.cpuOutput)
	return r
}

func part1() (*arcadeMachine, int) {
	a := newArcadeMachie()
	if err := a.run(); err != nil {
		fmt.Println("Part 1 failed with error:", err)
		return nil, 0
	}
	var blockCount int
	for _, t := range a.screen {
		if t == block {
			blockCount++
		}
	}
	fmt.Printf("Blocks: %d\n", blockCount)
	return a, blockCount
}

func part2() *arcadeMachine {
	a := newArcadeMachie()
	a.cpu.PatchMemory(0, []int{2})
	if err := a.run(); err != nil {
		fmt.Println("Part 2 failed with error:", err)
		return nil
	}
	fmt.Printf("Score: %d\n", a.score)
	return a
}
