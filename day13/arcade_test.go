package main

import (
	"testing"
)

func Test_arcadeMachine_cpuInput(t *testing.T) {
	tests := []struct {
		name    string
		ballX   int
		paddleX int
		want    int
	}{
		// TODO: Add test cases.
		{"same pos", 10, 10, 0},
		{"paddle left of ball", 10, 5, 1},
		{"paddle right of ball", 10, 15, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newArcadeMachie()
			a.ball.x = tt.ballX
			a.paddle.x = tt.paddleX
			if got := a.cpuInput(); got != tt.want {
				t.Errorf("arcadeMachine.cpuInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arcadeMachine_cpuOutput(t *testing.T) {
	tests := []struct {
		name          string
		params        []int
		wantBallPos   pos
		wantPaddlePos pos
		wantScore     int
	}{
		{"puzzle_example", []int{1, 2, 3, 6, 5, 4, -1, 0, 12345}, pos{6, 5}, pos{1, 2}, 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newArcadeMachie()
			for _, d := range tt.params {
				a.cpuOutput(d)
			}
			if a.ball != tt.wantBallPos {
				t.Errorf("ball pos = %v, want %v", a.ball, tt.wantBallPos)
			}
			if a.paddle != tt.wantPaddlePos {
				t.Errorf("paddle pos = %v, want %v", a.paddle, tt.wantPaddlePos)
			}
			if a.score != tt.wantScore {
				t.Errorf("score = %v, want %v", a.score, tt.wantScore)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	_, got := part1()
	want := 207
	if got != want {
		t.Errorf("part1() got = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	a := part2()
	want := 10247
	if a.score != want {
		t.Errorf("part1() got = %v, want %v", a.score, want)
	}
}
