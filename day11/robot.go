package main

import (
	"aoc2019/common/intcodecpu"
	"fmt"
	"strings"
)

type (
	pos struct {
		x, y int
	}
	robot struct {
		cpu         intcodecpu.CPU
		topLeft     pos
		bottomRight pos
		pos         pos
		dir         int // 0..3 = Up / Right / Down / Left
		out         int // 0..1 = Paint / Move
		grid        map[pos]int
	}
)

func (r *robot) cameraInput() int {
	if c, ok := r.grid[r.pos]; ok {
		return c
	}
	return 0 // Panel isn't painted yet, so it's black
}

func (r *robot) handleOutput(data int) {
	if r.out == 0 {
		r.paintPanel(data)
	} else {
		r.move(data)
	}
	r.out = 1 - r.out
}

func (r *robot) move(dir int) {
	if dir == 0 {
		r.dir--
		if r.dir < 0 {
			r.dir += 4
		}
	} else {
		r.dir++
	}
	r.dir = r.dir % 4
	switch r.dir {
	case 0:
		r.pos.y-- // Up
	case 1:
		r.pos.x++ // Right
	case 2:
		r.pos.y++ // Down
	case 3:
		r.pos.x-- // Left
	}
}

func (r *robot) run() error {
	r.pos = pos{0, 0}
	r.dir = 0
	r.out = 0
	return r.cpu.Run()
}

func (r *robot) paintPanel(color int) {
	r.grid[r.pos] = color
	if r.pos.x < r.topLeft.x {
		r.topLeft.x = r.pos.x
	} else if r.pos.x > r.bottomRight.x {
		r.bottomRight.x = r.pos.x
	}
	if r.pos.y < r.topLeft.y {
		r.topLeft.y = r.pos.y
	} else if r.pos.y > r.bottomRight.y {
		r.bottomRight.y = r.pos.y
	}
}

func (r *robot) getPanels() []string {
	height := r.bottomRight.y - r.topLeft.y + 1
	width := r.bottomRight.x - r.topLeft.x + 1
	panels := make([][]rune, height)
	for i := 0; i < height; i++ {
		panels[i] = []rune(strings.Repeat(" ", width))
	}
	for p, c := range r.grid {
		if c == 1 {
			panels[p.y][p.x] = '█'
		}
	}
	spanels := make([]string, height)
	for i := 0; i < height; i++ {
		spanels[i] = string(panels[i])
	}
	return spanels
}

func newHullPainterRobot() *robot {
	r := &robot{topLeft: pos{0, 0}, bottomRight: pos{0, 0}, pos: pos{0, 0}, dir: 0, out: 0, grid: make(map[pos]int)}
	r.cpu = *intcodecpu.NewProgramWithIO(program, r.cameraInput, r.handleOutput)
	return r
}

func part1() *robot {
	r := newHullPainterRobot()
	r.grid[pos{0, 0}] = 0 // Start on a black panel
	if err := r.run(); err != nil {
		fmt.Println("Part 1 failed with error:", err)
		return nil
	}
	fmt.Printf("Part 1 panels: %v (%v, %v -> %v, %v)\n", len(r.grid), r.topLeft.x, r.topLeft.y, r.bottomRight.x, r.bottomRight.y)
	return r
}

func part2() *robot {
	r := newHullPainterRobot()
	r.grid[pos{0, 0}] = 1 // Start on a white panel
	if err := r.run(); err != nil {
		fmt.Println("Part 2 failed with error:", err)
		return nil
	}
	fmt.Printf("Part 2 panels: %v (%v, %v -> %v, %v)\n", len(r.grid), r.topLeft.x, r.topLeft.y, r.bottomRight.x, r.bottomRight.y)
	for _, p := range r.getPanels() {
		fmt.Println(p)
	}
	return r
}
