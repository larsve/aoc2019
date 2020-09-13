package main

import (
	"aoc2019/common"
	"fmt"
	"reflect"
	"testing"
	"time"
)

var (
	smallTestGrid = []string{
		"█████",
		"█  ███",
		"█ █  █",
		"█ ● ██",
		"██████",
	}
	smallMultiBranchTestGrid = []string{
		"███████████████████████",
		"██ █ █ █ ██●██ █ █ █ ██",
		"█                     █",
		"██ ████████ ████████ ██",
		"███                 ███",
		"████ ██████ ██████ ████",
		"█████             █████",
		"██████ ████ ████ ██████",
		"███████         ███████",
		"████████ █ █ █ ████████",
		"███████████████████████",
	}
	multiBranchTestGrid = []string{
		"███████████████████████",
		"██                   ██",
		"█ █████████ █████████ █",
		"█ █ ██           ██ █ █",
		"█ █ █ █████ █████ █ █ █",
		"█ █ █ █ ██   ██ █ █ █ █",
		"█ █ █ █ █ █ █ █ █ █ █ █",
		"█          ●          █",
		"█ █ █ █ █ █ █ █ █ █ █ █",
		"█ █ █ █ ██   ██ █ █ █ █",
		"█ █ █ █████ █████ █ █ █",
		"█ █ ██           ██ █ █",
		"█ █████████ █████████ █",
		"██                   ██",
		"███████████████████████",
	}
	largeTestGrid = []string{
		"█████████████████████████████████████████",
		"█   █ █     █         █         █       █",
		"█ █ █ █ █ ███ ███████ █ ███ ███ ███ ███ █",
		"█ █   █ █     █ █   █   █   █ █ █   █ █ █",
		"█ ███ █ ███████ █ █ █████ ███ █ █ ███ █ █",
		"█   █ █ █       █ █       █   █ █   █   █",
		"█ █ █ █ █ █████ █ █ ███████ █ █ █ █ ███ █",
		"█ █ █ █ █ █   █ █ █ █   █   █ █ █ █   █ █",
		"███ █ █ █ █ █ ███ ███ █ ███ █ █ █████ ███",
		"█   █ █ █   █   █   █ █   █ █ █ █   █   █",
		"█ █████ █ █████ ███ █ ███ █ ███ █ █ ███ █",
		"█   █   █   █     █   █ █ █   █   █   █ █",
		"█ █ █ █████ █ ███ █████ █ █ █ ███████ █ █",
		"█ █   █   █ █ █   █     █ █ █       █ █ █",
		"█ █████ █ ███ █ ███████ █ █ ███████ █ █ █",
		"█   █   █     █ █     █ █ █ █ █     █   █",
		"███ █ █████████ █ ███ █ █ █ █ █ ███████ █",
		"█   █ █     █     █     █ █   █         █",
		"█ ███ █ █████ █████████ █ ███ ███████████",
		"█ █   █         █     █ █ █         █   █",
		"█ ███ █ █████████ ███ ███ █████████ █ ███",
		"█ █   █     █     █ █ █   █   █ █   █   █",
		"█ █ ███████ █ █████ ███ ███ █ █ █ █████ █",
		"█ █   █   █ █ █   █   █     █ █   █     █",
		"█ ███ █ ███ █ █ █ ███ █ █████ ███ █ █ ███",
		"█     █     █ █ █   █ █     █   █ █ █   █",
		"███████ █████ █ ███ █ █████ ███ █ █████ █",
		"█   █   █   █ █ █ █ █   █   █ █ █   █   █",
		"█ ███ █████ █ █ █ █ █ █ █ ███ █ ███ █ ███",
		"█ █   █     █   █   █ █ █     █ █ █ █   █",
		"█ █ ███ ███ █████ ███ ███ █████ █ █ ███ █",
		"█ █ █   █   █   █ █ █     █   █ █     █ █",
		"█ █ █ ███ █████ █ █ █ █████ █ █ █████ █ █",
		"█   █ █ █ █     █ █   █     █ █     █   █",
		"█ ███ █ █ █ ███ █ ███ █ █████ █ ███ ███ █",
		"█ █   █ █   █   █   █ █ █     █ █   █ █ █",
		"█ ███ █ ███ ███████ ███ █ ████░██ ███ █ █",
		"█ █   █   █ █       █   █ █   █   █     █",
		"█ █ ███ █ █ █ ███████ ███ █ █ █ ███ █████",
		"█   █   █   █         █     █   █      ●█",
		"█████████████████████████████████████████",
	}
)

func TestGetPossibleDirections(t *testing.T) {
	testGrid1 := []string{
		"█████",
		"██ ██",
		"█   █",
		"██ ██",
		"█████",
	}
	testGrid2 := []string{
		"███",
		"█ █",
		"█ █",
		"███",
	}
	tests := []struct {
		tg   []string
		sp   pos
		sd   direction
		want []direction
	}{
		{tg: testGrid1, sp: pos{2, 2}, want: []direction{north, south, west, east}},
		{tg: testGrid1, sp: pos{2, 2}, sd: north, want: []direction{south, west, east}},
		{tg: testGrid1, sp: pos{2, 2}, sd: east, want: []direction{north, south, west}},
		{tg: testGrid1, sp: pos{2, 2}, sd: south, want: []direction{north, west, east}},
		{tg: testGrid1, sp: pos{2, 2}, sd: west, want: []direction{north, south, east}},
		{tg: testGrid1, sp: pos{2, 1}, sd: south, want: []direction{}},

		{tg: testGrid2, sp: pos{1, 1}, sd: north, want: []direction{south}},
		{tg: testGrid2, sp: pos{1, 1}, sd: south, want: []direction{}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			g := make([][]rune, 5, 6)
			for i, r := range tt.tg {
				g[i] = []rune(r)
			}
			dl := getPossibleDirections(g, tt.sp, tt.sd)
			if len(dl) > 0 && len(tt.want) > 0 {
				common.Assert(t, reflect.DeepEqual(dl, tt.want), "getPossibleDirections path = %v, want %v", dl, tt.want)
			} else {
				common.Assert(t, len(dl) == len(tt.want), "getPossibleDirections path = %v, want %v", dl, tt.want)
			}
		})
	}

}

func TestFindLongestPath(t *testing.T) {
	tests := []struct {
		tg   []string
		sp   pos
		want int
	}{
		{tg: smallTestGrid, sp: pos{2, 3}, want: 4},
		{tg: smallMultiBranchTestGrid, sp: pos{11, 1}, want: 11},
		{tg: multiBranchTestGrid, sp: pos{11, 7}, want: 15},
		{tg: multiBranchTestGrid, sp: pos{11, 3}, want: 19},
		{tg: largeTestGrid, sp: pos{39, 39}, want: 310},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			g := make([][]rune, len(tt.tg), len(tt.tg[0]))
			for i, r := range tt.tg {
				g[i] = []rune(r)
			}
			start := time.Now()
			lp := findLongestPath(g, tt.sp)
			e := time.Since(start)
			t.Logf("Execution time: %v\n", e)
			common.Assert(t, lp == tt.want, "Longest path = %d, want %d", lp, tt.want)
		})
	}
}
