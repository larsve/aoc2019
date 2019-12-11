package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Diagnostic program--")
	fmt.Print("   Part 1 answer: ")
	fmt.Println(getDiagnosticCode(1))
	fmt.Print("   Part 2 answer: ")
	fmt.Println(getDiagnosticCode(5))
}
