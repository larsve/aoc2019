package main

import "fmt"

func main() {
	g := newWireGrid([]*wire{newWire(wire1), newWire(wire2)})
	fmt.Println(g.getShortestManhattanDistance())
	fmt.Println(g.getShortestStepPath())
}
