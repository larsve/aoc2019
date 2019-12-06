package intcodecpu

import "testing"

import "aoc2019/common"

func TestDecodeOpParams(t *testing.T) {
	test := func(p [4]int, e [6]int) {
		p1m, p1, p2m, p2, p3m, p3 := decodeOpParams(p[:])
		a := [6]int{p1m, p1, p2m, p2, p3m, p3}
		common.Assert(t, a == e, "Params %v did not generate the expected output: %v, got: %v", p, e, a)
	}
	test([4]int{0, 1, 2, 3}, [6]int{0, 1, 0, 2, 0, 3})
	test([4]int{100, 10, 20, 30}, [6]int{1, 10, 0, 20, 0, 30})
	test([4]int{1000, 15, 30, 45}, [6]int{0, 15, 1, 30, 0, 45})
	test([4]int{10000, 20, 40, 60}, [6]int{0, 20, 0, 40, 1, 60})
	test([4]int{12300, -1, -1, -1}, [6]int{3, -1, 2, -1, 1, -1})
}
