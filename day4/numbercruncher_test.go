package main

import (
	"testing"

	"aoc2019/common"
)

func TestSplit(t *testing.T) {
	test := func(actual, expected [6]int) {
		common.Assert(t, actual == expected, "Wrong value, expected %v, got%v", expected, actual)
	}
	test(split(123456), [6]int{1, 2, 3, 4, 5, 6})
	test(split(918273), [6]int{9, 1, 8, 2, 7, 3})
}

func TestHaveIncreasingDigits(t *testing.T) {
	test := func(n int, expected bool) {
		p := &password{digits: split(n)}
		common.Assert(t, p.haveIncreasingOrEqualDigits() == expected, "The number %v did not result in the expected result (%v)", n, expected)
	}
	test(111110, false)
	test(111111, true)
	test(123456, true)
	test(123321, false)
}

func TestHaveAdjacentDoubleDigitsPart1(t *testing.T) {
	test := func(n int, expected bool) {
		p := &password{digits: split(n)}
		common.Assert(t, p.haveAdjacentDoubleDigitsPart1() == expected, "The number %v did not result in the expected result (%v)", n, expected)
	}
	test(111111, true)
	test(123456, false)
	test(123321, true)
}

func TestHaveAdjacentDoubleDigitsPart2(t *testing.T) {
	test := func(n int, expected bool) {
		p := &password{digits: split(n)}
		common.Assert(t, p.haveAdjacentDoubleDigitsPart2() == expected, "The number %v did not result in the expected result (%v)", n, expected)
	}
	test(123456, false)
	test(112233, true)
	test(123444, false)
	test(123334, false)
	test(111122, true)
}

func TestGetPossiblePasswordsPart1(t *testing.T) {
	common.Assert(t, getPossiblePasswordsPart1(100000, 111110) == 0, "Test 1 failed")
	common.Assert(t, getPossiblePasswordsPart1(111110, 111112) == 2, "Test 2 failed")
	common.Assert(t, getPossiblePasswordsPart1(134564, 585159) == 1929, "Test 3 failed")
}

func TestGetPossiblePasswordsPart2(t *testing.T) {
	common.Assert(t, getPossiblePasswordsPart2(100000, 111110) == 0, "Test 1 failed")
	common.Assert(t, getPossiblePasswordsPart2(111110, 111122) == 1, "Test 2 failed")
	common.Assert(t, getPossiblePasswordsPart2(134564, 585159) == 1306, "Test 3 failed")
}

func BenchmarkGetPossiblePasswordsPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = getPossiblePasswordsPart1(134564, 585159)
	}
}

func BenchmarkGetPossiblePasswordsPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = getPossiblePasswordsPart2(134564, 585159)
	}
}
