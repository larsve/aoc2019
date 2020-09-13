package main

import (
	"aoc2019/common/intcodecpu"
	"fmt"
	"strings"
)

type (
	pos struct {
		x int
		y int
	}
	droid struct {
		cpu          *intcodecpu.CPU
		breadCrumbs  bool
		haltOnOxySys bool
		pos          pos
		dir          direction
		minX         int
		maxX         int
		minY         int
		maxY         int
		spaceMap     map[pos]rune
		oxySysPos    pos
	}
)

const (
	start   = '⌂'
	wall    = '█'
	oxygen  = '◙'
	space   = ' '
	unknown = '■'
	deadEnd = '▫'
)

var (
	startPos = pos{0, 0}
)

func (d *droid) generateGrid() (grid [][]rune, oxygenStation pos) {
	width := d.maxX - d.minX + 1
	height := d.maxY - d.minY + 1
	xOfs := -d.minX
	yOfs := -d.minY
	oxygenStation = pos{d.oxySysPos.x + xOfs, d.oxySysPos.y + yOfs}
	grid = make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = []rune(strings.Repeat(string(unknown), width))
	}
	for p, c := range d.spaceMap {
		grid[yOfs+p.y][xOfs+p.x] = c
	}
	return
}

func (d *droid) getMinimumNumberOfMovements() int {
	var steps int
	for _, c := range d.spaceMap {
		switch c {
		case oxygen, north.breadCrumb(), east.breadCrumb(), south.breadCrumb(), west.breadCrumb():
			steps++
		}
	}
	return steps
}

func (d *droid) getSpaceMap() string {
	grid, _ := d.generateGrid()
	height := len(grid)
	spanels := make([]string, height)
	for i := 0; i < height; i++ {
		spanels[i] = strings.ReplaceAll(string(grid[i]), string(unknown), string(wall))
	}
	return strings.Join(spanels, "\n")
}

func (d *droid) setMinMaxPos(p pos) {
	if d.minX > p.x {
		d.minX = p.x
	} else if d.maxX < p.x {
		d.maxX = p.x
	}
	if d.minY > p.y {
		d.minY = p.y
	} else if d.maxY < p.y {
		d.maxY = p.y
	}
}

func (d *droid) stdIn() int {
	//fmt.Printf("You are at %v\tGo %s.. ", d.pos, d.dir)
	return int(d.dir)
}

func (d *droid) stdOut(data int) {
	pos := d.dir.newPos(d.pos)
	d.setMinMaxPos(pos)
	switch data {
	case 0:
		// Wall, could not move
		d.hitWall(pos)
	case 1:
		// Moved
		d.movedOk(pos)
	case 2:
		// Moved and found oxygen system
		//fmt.Println("You have found the oxygen system!")
		d.pos = pos
		d.oxySysPos = pos
		if d.breadCrumbs {
			d.spaceMap[pos] = oxygen
		} else {
			d.spaceMap[pos] = space
		}
		if d.haltOnOxySys {
			d.cpu.Halt()
		}
	}
}

func (d *droid) hitWall(p pos) {
	//fmt.Printf("There is a wall there, you can't go %s\n", d.dir)
	d.spaceMap[p] = wall
	d.dir = d.dir.turnRight()
}

func (d *droid) movedOk(p pos) {
	//fmt.Println("OK")
	from := d.pos
	d.pos = p
	if d.breadCrumbs {
		if crumb, ok := d.spaceMap[p]; !ok {
			d.spaceMap[p] = d.dir.breadCrumb()
		} else {
			//fmt.Printf("You have been here before, you left a %c bread crumb..\n", crumb)
			if crumb == d.dir.opposite().breadCrumb() {
				d.spaceMap[p] = deadEnd // Reversing in my own footsteps...
			}
			d.spaceMap[from] = deadEnd
		}
	} else {
		d.spaceMap[p] = space
	}
	if d.pos == startPos {
		d.cpu.Halt()
	}
	d.dir = d.dir.turnLeft()
}

func newDroid() *droid {
	d := &droid{
		breadCrumbs:  true,
		haltOnOxySys: true,
		dir:          north,
		spaceMap:     make(map[pos]rune),
	}
	d.cpu = intcodecpu.NewProgramWithIO(program, d.stdIn, d.stdOut)
	d.spaceMap[d.pos] = start
	return d
}

func part1() *droid {
	d := newDroid()
	if err := d.cpu.Run(); err != nil {
		fmt.Printf("Faild to run droid program, error: %v\n", err)
		return nil
	}
	fmt.Println(d.getSpaceMap())
	fmt.Printf("Mimimum number of steps: %d\n", d.getMinimumNumberOfMovements())
	return d
}

func part2() int {
	d := newDroid()
	d.breadCrumbs = false
	d.haltOnOxySys = false
	if err := d.cpu.Run(); err != nil {
		fmt.Printf("Faild to run droid program, error: %v\n", err)
		return -1
	}
	g, osp := d.generateGrid()
	maxLen := findLongestPath(g, osp)
	fmt.Printf("Maximum length from oxygen station: %d\n", maxLen)
	return maxLen
}
