package main

import (
	"aoc2019/common"
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

var (
	example1 = `10 ORE => 10 A
	1 ORE => 1 B
	7 A, 1 B => 1 C
	7 A, 1 C => 1 D
	7 A, 1 D => 1 E
	7 A, 1 E => 1 FUEL`
	example2 = `9 ORE => 2 A
	8 ORE => 3 B
	7 ORE => 5 C
	3 A, 4 B => 1 AB
	5 B, 7 C => 1 BC
	4 C, 1 A => 1 CA
	2 AB, 3 BC, 4 CA => 1 FUEL`
	example3 = `157 ORE => 5 NZVS
	165 ORE => 6 DCFZ
	44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
	12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
	179 ORE => 7 PSHF
	177 ORE => 5 HKGWZ
	7 DCFZ, 7 PSHF => 2 XJWVT
	165 ORE => 2 GPVTF
	3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`
	example4 = `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
	17 NVRVD, 3 JNWZP => 8 VPVL
	53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
	22 VJHF, 37 MNCFX => 5 FWMGM
	139 ORE => 4 NVRVD
	144 ORE => 7 JNWZP
	5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
	5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
	145 ORE => 6 MNCFX
	1 NVRVD => 8 CXFTF
	1 VJHF, 6 MNCFX => 4 RFSQX
	176 ORE => 6 VJHF`
	example5 = `171 ORE => 8 CNZTR
	7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
	114 ORE => 4 BHXH
	14 VRPVC => 6 BMBT
	6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
	6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
	15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
	13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
	5 BMBT => 4 WPTQ
	189 ORE => 9 KTJDG
	1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
	12 VRPVC, 27 CNZTR => 2 XDBXC
	15 KTJDG, 12 BHXH => 5 XCVML
	3 BHXH, 2 VRPVC => 7 MZWV
	121 ORE => 7 VRPVC
	7 XCVML => 6 RJRHP
	5 BHXH, 4 VRPVC => 5 LTCX`
)

func TestReactionParse(t *testing.T) {
	tt := []struct {
		data           string
		wantedInput    []reactionChemical
		wantedOutput   reactionChemical
		wantedError    error
		wantedErrorMsg string
	}{
		{
			data:         "8 ORE => 3 B",
			wantedInput:  []reactionChemical{{qty: 8, chm: "ORE"}},
			wantedOutput: reactionChemical{qty: 3, chm: "B"},
			wantedError:  nil,
		},
		{
			data:         "1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP",
			wantedInput:  []reactionChemical{{qty: 1, chm: "MZWV"}, {qty: 17, chm: "XDBXC"}, {qty: 3, chm: "XCVML"}},
			wantedOutput: reactionChemical{qty: 2, chm: "XMNCP"},
			wantedError:  nil,
		},
		{
			data:        "one_token",
			wantedError: errorInvalidReactionInput,
		},
		{
			data:           "one ORE",
			wantedErrorMsg: `strconv.Atoi: parsing "one": invalid syntax`,
		},
		{
			data:        "10 ORE",
			wantedError: errorInvalidReactionInput,
		},
	}
	for _, tc := range tt {
		t.Run(tc.data, func(t *testing.T) {
			r := &reaction{}
			err := r.parse(tc.data)
			if tc.wantedError == nil && tc.wantedErrorMsg == "" {
				common.Assert(t, err == nil, "Parse failed, error %v", err)
				common.Assert(t, reflect.DeepEqual(r.input, tc.wantedInput), "Parse failed, not the expected input, got %v but expected %v", r.input, tc.wantedInput)
				common.Assert(t, r.output == tc.wantedOutput, "Parse failed, not the expected output, got %v but expected %v", r.output, tc.wantedOutput)
			} else if tc.wantedError != nil {
				common.Assert(t, errors.Is(err, tc.wantedError), "Parse failed with unexpected error, error %v", err)
			} else {
				common.Assert(t, strings.Contains(err.Error(), tc.wantedErrorMsg), "Parse failed with unexpected error, error %v", err)
			}
		})
	}
}

func TestNewNanoFactory(t *testing.T) {
	tt := []struct {
		name                  string
		recipe                string
		wantedReactionCnt     int
		wantedFuelIngredients []reactionChemical
	}{
		{
			name:                  "Example 1",
			recipe:                example1,
			wantedReactionCnt:     6,
			wantedFuelIngredients: []reactionChemical{{qty: 7, chm: "A"}, {qty: 1, chm: "E"}},
		},
		{
			name:                  "Example 2",
			recipe:                example2,
			wantedReactionCnt:     7,
			wantedFuelIngredients: []reactionChemical{{qty: 2, chm: "AB"}, {qty: 3, chm: "BC"}, {qty: 4, chm: "CA"}},
		},
		{
			name:                  "Example 3",
			recipe:                example3,
			wantedReactionCnt:     9,
			wantedFuelIngredients: []reactionChemical{{qty: 44, chm: "XJWVT"}, {qty: 5, chm: "KHKGT"}, {qty: 1, chm: "QDVJ"}, {qty: 29, chm: "NZVS"}, {qty: 9, chm: "GPVTF"}, {qty: 48, chm: "HKGWZ"}},
		},
		{
			name:                  "Example 4",
			recipe:                example4,
			wantedReactionCnt:     12,
			wantedFuelIngredients: []reactionChemical{{qty: 53, chm: "STKFG"}, {qty: 6, chm: "MNCFX"}, {qty: 46, chm: "VJHF"}, {qty: 81, chm: "HVMC"}, {qty: 68, chm: "CXFTF"}, {qty: 25, chm: "GNMV"}},
		},
		{
			name:                  "Example 5",
			recipe:                example5,
			wantedReactionCnt:     17,
			wantedFuelIngredients: []reactionChemical{{qty: 6, chm: "BHXH"}, {qty: 18, chm: "KTJDG"}, {qty: 12, chm: "WPTQ"}, {qty: 7, chm: "PLWSL"}, {qty: 31, chm: "FHTLT"}, {qty: 37, chm: "ZDVW"}},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			nf, err := newNanoFactory(bytes.NewBufferString(tc.recipe))
			common.Assert(t, err == nil, "Failed to create NanoFactory, error: %v", err)
			common.Assert(t, nf != nil, "Failed to create NanoFactory, returned nil")
			common.Assert(t, len(nf.reactions) == tc.wantedReactionCnt, "Not expected number of reactions, got %d, but expected %d", len(nf.reactions), tc.wantedReactionCnt)
			r, ok := nf.reactions["FUEL"]
			common.Assert(t, ok, "No fuel found")
			common.Assert(t, reflect.DeepEqual(r.input, tc.wantedFuelIngredients), "Not the expected ingredients, got %v, but expected %v", r.input, tc.wantedFuelIngredients)
		})
	}
}

func TestNewNanoFactoryInStock(t *testing.T) {
	nf := &nanoFactory{
		produced: make(map[chemical]int),
		consumed: make(map[chemical]int),
	}
	q := nf.inStock(ore)
	common.Assert(t, q == 0, "Maps are empty, inStock should return zero, got %d", q)

	nf.produced[fuel] = 10
	q = nf.inStock(fuel)
	common.Assert(t, q == 10, "We have produced 10 but not consumed any, inStock should return 10, got %d", q)

	nf.consumed[fuel] = 7
	q = nf.inStock(fuel)
	common.Assert(t, q == 3, "We have produced 10 and consumed 7, inStock should return 3, got %d", q)
}

func TestNewNanoFactoryNeed(t *testing.T) {
	tt := []struct {
		name   string
		recipe string
		oreCnt int
	}{
		{name: "Example 1", recipe: example1, oreCnt: 31},
		{name: "Example 2", recipe: example2, oreCnt: 165},
		{name: "Example 3", recipe: example3, oreCnt: 13312},
		{name: "Example 4", recipe: example4, oreCnt: 180697},
		{name: "Example 5", recipe: example5, oreCnt: 2210736},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			nf, err := newNanoFactory(bytes.NewBufferString(tc.recipe))
			common.Assert(t, err == nil, "error %v", err)
			err = nf.need(1, fuel)
			common.Assert(t, err == nil, "error %v", err)
			common.Assert(t, nf.consumed[ore] == tc.oreCnt, "got %d, expected %d", nf.consumed[ore], tc.oreCnt)
		})
	}
}

func TestNewNanoFactoryFuelForOre(t *testing.T) {
	tt := []struct {
		name    string
		recipe  string
		fuelCnt int
	}{
		{name: "Example 3", recipe: example3, fuelCnt: 82892753},
		{name: "Example 4", recipe: example4, fuelCnt: 5586022},
		{name: "Example 5", recipe: example5, fuelCnt: 460664},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			nf, err := newNanoFactory(bytes.NewBufferString(tc.recipe))
			common.Assert(t, err == nil, "error %v", err)
			fc, err := nf.fuelForOre(oreSupply)
			common.Assert(t, err == nil, "error %v", err)
			common.Assert(t, fc == tc.fuelCnt, "got %d, expected %d", fc, tc.fuelCnt)
		})
	}
}

func TestIntRoundUp(t *testing.T) {
	d := 3
	a := []int{0, 1, 1, 1, 2, 2, 2, 3, 3, 3}
	for i, j := range a {
		k := (i + d - 1) / d
		common.Assert(t, k == j, "got %d expected %d", k, j)
	}

	d = 2
	a = []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5}
	for i, j := range a {
		k := (i + d - 1) / d
		common.Assert(t, k == j, "got %d expected %d", k, j)
	}
}

func TestPart1(t *testing.T) {
	nf := part1("input")
	common.Assert(t, nf != nil, "No factory returned")
	common.Assert(t, nf.consumed[ore] == 374457, "got %d, wanted 374457", nf.consumed[ore])
}

func TestPart2(t *testing.T) {
	nf, fuel := part2("input")
	common.Assert(t, nf != nil, "No factory returned")
	common.Assert(t, fuel == 3568888, "got %d, wanted 3568888", fuel)
}
