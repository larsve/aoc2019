package main

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"testing"
	"time"
)

func TestSignalString(t *testing.T) {
	tests := []struct {
		input signal
		want  string
	}{
		{input: []byte{}, want: ""},
		{input: []byte{1}, want: "1"},
		{input: []byte{1, 2, 3, 4, 5, 6, 7, 8}, want: "12345678"},
		{input: []byte{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5}, want: "80871224585914546619083218645595"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.input.String(); got != tt.want {
				t.Errorf("signal.String() got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalValue(t *testing.T) {
	tests := []struct {
		input signal
		want  string
	}{
		{input: []byte{}, want: ""},
		{input: []byte{1}, want: "1"},
		{input: []byte{1, 2, 3, 4, 5, 6, 7, 8}, want: "12345678"},
		{input: []byte{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5}, want: "80871224"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.input.Value(); got != tt.want {
				t.Errorf("signal.Value() got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToSignal(t *testing.T) {
	tests := []struct {
		input string
		want  signal
	}{
		{input: "12345678", want: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
		{input: "80871224585914546619083218645595", want: []byte{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5}},
		{input: "19617804207202209144916044189917", want: []byte{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7}},
		{input: "69317163492948606335995924319873", want: []byte{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3}},
		{input: " @Jh", want: []byte{0, 6, 6, 6}},
		{input: "☻", want: []byte{1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := stos(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stos(%s) got %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFft(t *testing.T) {
	tests := []struct {
		input  signal
		phases int
		want   signal
	}{
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 1, want: []byte{4, 8, 2, 2, 6, 1, 5, 8}},
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 2, want: []byte{3, 4, 0, 4, 0, 4, 3, 8}},
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 3, want: []byte{0, 3, 4, 1, 5, 5, 1, 8}},
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 4, want: []byte{0, 1, 0, 2, 9, 4, 9, 8}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test#%d", i), func(t *testing.T) {
			signal := tt.input
			for j := 0; j < tt.phases; j++ {
				signal = fft(signal)
			}
			//got := signal*[:8]
			if !reflect.DeepEqual(signal, tt.want) {
				t.Errorf("fft got %v after %d phases, want %v", signal, tt.phases, tt.want)
			}
		})
	}
}

func TestFftLoop(t *testing.T) {
	tests := []struct {
		input  signal
		phases int
		want   string
	}{
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 1, want: "48226158"},
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 2, want: "34040438"},
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 3, want: "03415518"},
		{input: signal{1, 2, 3, 4, 5, 6, 7, 8}, phases: 4, want: "01029498"},
		{input: signal{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5}, phases: 100, want: "24176176"},
		{input: signal{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7}, phases: 100, want: "73745418"},
		{input: signal{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3}, phases: 100, want: "52432133"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test#%d", i), func(t *testing.T) {
			got := fftLoop(tt.input, tt.phases).Value()
			if got != tt.want {
				t.Errorf("fft got %s after %d phases, want %s", got, tt.phases, tt.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	want := "82435530"
	got := part1().Value()
	if got != want {
		t.Errorf("part1 got %s, want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	timeout := time.Second * 10
	tot := time.AfterFunc(timeout, func() {
		debug.SetTraceback("all")
		panic(fmt.Sprintf("test timed out after %v", timeout))
	})
	defer tot.Stop()
	want := "?"
	got := part2().Value()
	if got != want {
		t.Errorf("part2 got %s, want %s", got, want)
	}
}
