package main

import "fmt"

func main() {
	p := newProgram(code[:])
	p.debug = true
	// Restore the "1202 program alert code"
	p.setAlarmCode(1202)
	fmt.Println(p.run())

	fmt.Println(getNounAndVerbFor(19690720))
}
