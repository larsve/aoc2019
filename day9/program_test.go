package main

import (
	"testing"

	"aoc2019/common"
)

func TestBoostKeycode(t *testing.T) {
	c, e := getBoostKeycode(false)
	common.Assert(t, e == nil, "TestBoostKeycode falied with error: %v", e)
	common.Assert(t, c == 3380552333, "TestBoostKeycode did not generate expected BOOST keycode, got: %v", c)
}

func TestGetDistressCoorinates(t *testing.T) {
	c, e := getDistressCoorinates(false)
	common.Assert(t, e == nil, "TestGetDistressCoorinates falied with error: %v", e)
	common.Assert(t, c == 78831, "TestGetDistressCoorinates did not generate expected coorinates, got: %v", c)
}
