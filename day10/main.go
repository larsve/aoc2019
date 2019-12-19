package main

import "fmt"

func main() {
	fmt.Println("-- Monitoring station --")
	m, e := newMapFromFile(42, 42, "day10/input")
	if e != nil {
		fmt.Println("Failed to read map:", e)
		return
	}
	fmt.Println("Part 1:")
	m.getBestPlacement()
	fmt.Println("Part 2:")
	m.part2Answer()
}
