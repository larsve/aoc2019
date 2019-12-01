package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRequiredFuel(t *testing.T) {
	tt := []struct {
		mass int64 // a module mass
		fuel int64 // expected fuel
	}{
		{mass: 12, fuel: 2},
		{mass: 14, fuel: 2},
		{mass: 1969, fuel: 654},
		{mass: 100756, fuel: 33583},
		{mass: 29, fuel: 7},
		{mass: 30, fuel: 8},
		{mass: 0, fuel: 0},
		{mass: 1, fuel: 0},
		{mass: 2, fuel: 0},
		{mass: 8, fuel: 0},
		{mass: 9, fuel: 1},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Mass %v", tc.mass), func(t *testing.T) {
			rf := requiredFuel(tc.mass)
			assert(t, rf == tc.fuel, "wrong fuel for mass %v, expected %v but got %v", tc.mass, tc.fuel, rf)
		})

	}
}

func BenchmarkRequiredFuel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		requiredFuel(42)
	}
}

func TestTotalRequiredFuel(t *testing.T) {
	tt := []struct {
		masses []int64 // module mass list to calculate required fuel for
		fuel   int64   // expected fuel
	}{
		{masses: []int64{0}, fuel: 0},
		{masses: []int64{12, 14, 1969, 100756}, fuel: 34241},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("totalRequiredFuel test #%v", i), func(t *testing.T) {
			rf := totalRequiredFuel(tc.masses)
			assert(t, rf == tc.fuel, "wrong fuel for test %v, expected %v but got %v", i, tc.fuel, rf)
		})

	}
}

func BenchmarkTotalRequiredFuel(b *testing.B) {
	masses := []int64{12, 14, 1969, 100756}
	for n := 0; n < b.N; n++ {
		totalRequiredFuel(masses)
	}
}

func TestMassCompensatedFuel(t *testing.T) {
	tt := []struct {
		mass int64 // starting mass
		fuel int64 // expected compensated fuel
	}{
		{mass: 14, fuel: 2},
		{mass: 1969, fuel: 966},
		{mass: 100756, fuel: 50346},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("massCompensatedFuel %v", tc.mass), func(t *testing.T) {
			rf := massCompensatedFuel(requiredFuel(tc.mass))
			assert(t, rf == tc.fuel, "wrong fuel for mass %v, expected %v but got %v", tc.mass, tc.fuel, rf)
		})

	}
}

func benchmarkMassCompensatedFuel(mass int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		massCompensatedFuel(requiredFuel(mass))
	}
}
func BenchmarkMassCompensatedFuel1(b *testing.B)       { benchmarkMassCompensatedFuel(1, b) }
func BenchmarkMassCompensatedFuel10(b *testing.B)      { benchmarkMassCompensatedFuel(10, b) }
func BenchmarkMassCompensatedFuel100(b *testing.B)     { benchmarkMassCompensatedFuel(100, b) }
func BenchmarkMassCompensatedFuel1000(b *testing.B)    { benchmarkMassCompensatedFuel(1000, b) }
func BenchmarkMassCompensatedFuel10000(b *testing.B)   { benchmarkMassCompensatedFuel(10000, b) }
func BenchmarkMassCompensatedFuel100000(b *testing.B)  { benchmarkMassCompensatedFuel(100000, b) }
func BenchmarkMassCompensatedFuel1000000(b *testing.B) { benchmarkMassCompensatedFuel(1000000, b) }

func TestTotalMassCompensatedRequiredFuel(t *testing.T) {
	tt := []struct {
		masses []int64 // module masses
		fuel   int64   // expected fuel
	}{
		{masses: []int64{0}, fuel: 0},
		{masses: []int64{14}, fuel: 2},
		{masses: []int64{14, 1969, 100756}, fuel: 2 + 966 + 50346},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("totalMassCompensatedFuel test #%v", i), func(t *testing.T) {
			rf := totalMassCompensatedFuel(tc.masses)
			assert(t, rf == tc.fuel, "wrong fuel for test %v, expected %v but got %v", i, tc.fuel, rf)
		})

	}
}

func BenchmarkTotalMassCompensatedRequiredFuel(b *testing.B) {
	masses := []int64{12, 14, 1969, 100756}
	for n := 0; n < b.N; n++ {
		totalMassCompensatedFuel(masses)
	}
}

func BenchmarkFirstAnswer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		totalRequiredFuel(shipModuleMasses)
	}
}

func BenchmarkSecondAnswer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		totalMassCompensatedFuel(shipModuleMasses)
	}
}

func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}
