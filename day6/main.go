package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("input")
	if e != nil {
		fmt.Println("Failed to open file, error:", e)
		return
	}
	defer f.Close()
	m, e := newMap(f)
	if e != nil {
		fmt.Println("Failed to read orbital map, error:", e)
		return
	}
	fmt.Printf("Checksum: %v", m.orbitCountChecksum)
	fmt.Printf("YOU => SAN distance: %v", m.getDistance("YOU", "SAN"))
}
