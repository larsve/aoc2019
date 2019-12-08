package main

import (
	"bytes"
	"os"
	"testing"

	"aoc2019/common"
)

const (
	part1ExampleMap = "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	part2ExampleMap = "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN\n"
)

func TestOrbitalMap(t *testing.T) {
	m, e := newMap(bytes.NewBufferString(part1ExampleMap))
	common.Assert(t, e == nil, "Failed to create test map, ended with error: %v", e)
	common.Assert(t, m.orbitCountChecksum == 42, "wrong checksum: %v", m.orbitCountChecksum)

	f, e := os.Open("input")
	common.Assert(t, e == nil, "Unable to open input map file: %v", e)
	defer f.Close()
	m, e = newMap(f)
	common.Assert(t, e == nil, "Failed to create map, ended with error: %v", e)
	common.Assert(t, m.orbitCountChecksum == 417916, "wrong checksum: %v", m.orbitCountChecksum)
}

func TestDistance(t *testing.T) {
	m, e := newMap(bytes.NewBufferString(part2ExampleMap))
	common.Assert(t, e == nil, "Failed to create test map, ended with error: %v", e)
	d := m.getDistance("YOU", "SAN")
	common.Assert(t, d == 4, "Not the expected distance: %v, expected 4", d)

	f, e := os.Open("input")
	common.Assert(t, e == nil, "Unable to open input map file: %v", e)
	defer f.Close()
	m, e = newMap(f)
	common.Assert(t, e == nil, "Failed to create map, ended with error: %v", e)
	d = m.getDistance("YOU", "SAN")
	common.Assert(t, d == 523, "wrong distance: %v", d)
}
