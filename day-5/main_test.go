package main

import (
	"testing"
)

func Test_solvePart1(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{
			name:       "example",
			args:       args{"test-input"},
			wantAnswer: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := solvePart1(tt.args.inputFile); gotAnswer != tt.wantAnswer {
				t.Errorf("solvePart1() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer int
	}{
		{
			name:       "example",
			args:       args{"test-input"},
			wantAnswer: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := solvePart2(tt.args.inputFile); gotAnswer != tt.wantAnswer {
				t.Errorf("solvePart2() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
