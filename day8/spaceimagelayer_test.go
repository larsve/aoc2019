package main

import "testing"

import "aoc2019/common"

func TestNewLayer(t *testing.T) {
	rawData := []byte{'0', '1', '2', '3', '4', '5', '5', '5', '4', '3', '2', '1', '0'}
	l := newLayer(rawData)
	common.Assert(t, len(l.rawData) == len(rawData), "Layer raw data is not of the expected length, expected %v, but got %v", len(rawData), len(l.rawData))
	common.Assert(t, l.distibution[4] == 2, "Wrong distribution count for '4', expected 2, but got %v", l.distibution[4])
	common.Assert(t, l.distibution[5] == 3, "Wrong distribution count for '5', expected 3, but got %v", l.distibution[5])
}
