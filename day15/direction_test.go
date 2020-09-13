package main

import "testing"

func TestDirection(t *testing.T) {
	tests := []struct {
		d    direction
		wInt int
		wStr string
		wBC  rune
		wOD  direction
		wTR  direction
		wTL  direction
		wPos pos
	}{
		{d: north, wInt: 1, wStr: "North", wBC: '▲', wOD: south, wTR: east, wTL: west, wPos: pos{0, -1}},
		{d: east, wInt: 4, wStr: "East", wBC: '►', wOD: west, wTR: south, wTL: north, wPos: pos{1, 0}},
		{d: south, wInt: 2, wStr: "South", wBC: '▼', wOD: north, wTR: west, wTL: east, wPos: pos{0, 1}},
		{d: west, wInt: 3, wStr: "West", wBC: '◄', wOD: east, wTR: north, wTL: south, wPos: pos{-1, 0}},
	}
	startPos := pos{0, 0}
	for _, tt := range tests {
		t.Run(tt.d.String(), func(t *testing.T) {
			if got := int(tt.d); got != tt.wInt {
				t.Errorf("direction = %v, want %v", got, tt.wInt)
			}
			if got := tt.d.String(); got != tt.wStr {
				t.Errorf("direction.String() = %v, want %v", got, tt.wStr)
			}
			if got := tt.d.breadCrumb(); got != tt.wBC {
				t.Errorf("direction.breadCrumb() = %v, want %v", got, tt.wBC)
			}
			if got := tt.d.opposite(); got != tt.wOD {
				t.Errorf("direction.opposite() = %v, want %v", got, tt.wOD)
			}
			if got := tt.d.turnRight(); got != tt.wTR {
				t.Errorf("direction.turnRight() = %v, want %v", got, tt.wTR)
			}
			if got := tt.d.turnLeft(); got != tt.wTL {
				t.Errorf("direction.turnLeft() = %v, want %v", got, tt.wTL)
			}
			if got := tt.d.newPos(startPos); got != tt.wPos {
				t.Errorf("direction.newPos() = %v, want %v", got, tt.wPos)
			}
		})
	}
}
